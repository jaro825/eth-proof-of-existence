package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

type Contract struct {
	Address      string `json:"address,omitempty"`
	OwnerAddress string `json:"owner_address,omitempty"`
	TxID         string `json:"tx_id,omitempty"`
	ProofsStored uint64 `json:"proofs_stored,omitempty"`
}

func ActionContractDeploy(ctx *cli.Context) error {
	privateKeyHex := getStringParamFromContext(ctx, "private_key")
	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("HexToECDSA error: %w", err)
	}

	walletAddressHex := getStringParamFromContext(ctx, "address")
	address := common.HexToAddress(walletAddressHex)

	nodeURL := getStringParamFromContext(ctx, "node_url")
	backend, err := NewNetworkBackend(nodeURL)
	if err != nil {
		return fmt.Errorf("NewNetworkBackend error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()

	nonce, err := backend.PendingNonceAt(reqCtx, address)
	if err != nil {
		return fmt.Errorf("PendingNonceAt error: %w", err)
	}

	gasPrice, err := backend.SuggestGasPrice(reqCtx)
	if err != nil {
		return fmt.Errorf("SuggestGasPrice error: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, backend.ChainId)
	if err != nil {
		return fmt.Errorf("NewKeyedTransactorWithChainID error: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	contractAddress, tx, _, err := DeployStore(auth, backend)
	if err != nil {
		return fmt.Errorf("DeployStore error: %w", err)
	}

	contract := &Contract{
		Address: contractAddress.String(),
		TxID:    tx.Hash().Hex(),
	}

	fmt.Println("--------------------")
	fmt.Println("contract deployed:")
	prettyPrint(contract)

	return nil
}

func ActionContractStatus(ctx *cli.Context) error {
	contractAddressHex := getStringParamFromContext(ctx, "address")
	address := common.HexToAddress(contractAddressHex)

	nodeURL := getStringParamFromContext(ctx, "node_url")
	backend, err := NewNetworkBackend(nodeURL)
	if err != nil {
		return fmt.Errorf("NewNetworkBackend error: %w", err)
	}

	instance, err := NewStore(address, backend)
	if err != nil {
		return fmt.Errorf("NewStore error: %w", err)
	}

	ownerAddress, err := instance.Owner(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("Owner error: %w", err)
	}

	size, err := instance.GetSize(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("GetSize error: %w", err)
	}

	contract := &Contract{
		Address:      address.String(),
		OwnerAddress: ownerAddress.String(),
		ProofsStored: size.Uint64(),
	}

	fmt.Println("--------------------")
	fmt.Println("contract status:")
	prettyPrint(contract)

	return nil
}
