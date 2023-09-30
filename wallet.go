package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/common/hexutil"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKeyHex string `json:"private_key_hex"`
	PublicKeyHex  string `json:"public_key_hex"`
	AddressHex    string `json:"address_hex"`
}

func CreateNewWallet() (*Wallet, error) {
	privateKey, err := ethcrypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := ethcrypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")

	}

	publicKeyBytes := ethcrypto.FromECDSAPub(publicKeyECDSA)

	addressHex := ethcrypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &Wallet{
		PrivateKeyHex: hexutil.Encode(privateKeyBytes),
		PublicKeyHex:  hexutil.Encode(publicKeyBytes),
		AddressHex:    addressHex,
	}, nil

}

func ActionWalletCreate(_ *cli.Context) error {
	wallet, err := CreateNewWallet()
	if err != nil {
		return fmt.Errorf("could not create wallet: %w", err)
	}

	fmt.Println("wallet created:")

	b, err := json.MarshalIndent(wallet, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal wallet: %w", err)
	}

	fmt.Println("--------------------")
	fmt.Printf("%s\n", string(b))

	return nil
}

func ActionWalletBalance(ctx *cli.Context) error {
	addressHex := ctx.Value("address")

	account := common.HexToAddress(addressHex.(string))

	nodeURL := ctx.Value("node_url")

	backend, err := NewNetworkBackend(nodeURL.(string))
	if err != nil {
		return fmt.Errorf("could not create network connection to backend: %w", err)
	}

	requestCtx, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()

	balance, err := backend.BalanceAt(requestCtx, account, nil)
	if err != nil {
		return fmt.Errorf("could not get balance: %w", err)
	}

	fmt.Println("--------------------")
	fmt.Printf("the balance of %s is %d wei\n", addressHex, balance)

	return nil
}