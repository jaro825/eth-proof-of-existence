package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func ActionFileAdd(ctx *cli.Context) error {
	privateKeyHex := ctx.Value("private_key")
	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex.(string))
	if err != nil {
		return fmt.Errorf("HexToECDSA error: %w", err)
	}

	contractAddressHex := ctx.Value("contract_address")
	contractAddress := common.HexToAddress(contractAddressHex.(string))

	ownerAddressHex := ctx.Value("contract_owner")
	ownerAddress := common.HexToAddress(ownerAddressHex.(string))

	filePath := ctx.Value("file")
	hash := fileHash(filePath.(string))
	fileBasename := filepath.Base(filePath.(string))
	fileUUID := generateUUID(fileBasename)

	nodeURL := ctx.Value("node_url")

	backend, err := NewNetworkBackend(nodeURL.(string))
	if err != nil {
		return fmt.Errorf("NewNetworkBackend error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, backend.ChainId)
	if err != nil {
		return fmt.Errorf("NewKeyedTransactorWithChainID error: %w", err)
	}

	nonce, err := backend.PendingNonceAt(reqCtx, ownerAddress)
	if err != nil {
		return fmt.Errorf("PendingNonceAt error: %w", err)
	}

	gasPrice, err := backend.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SuggestGasPrice error: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	instance, err := NewStore(contractAddress, backend)
	if err != nil {
		return fmt.Errorf("NewStore error: %w", err)
	}

	key := [16]byte(fileUUID[:])

	encodedFileName := make([]byte, base64.StdEncoding.EncodedLen(len(fileBasename)))

	base64.StdEncoding.Encode(encodedFileName, []byte(fileBasename))

	created := uint32(time.Now().Unix())

	add, err := instance.Add(auth, key, encodedFileName, hash, created)
	if err != nil {
		return fmt.Errorf("Add error: %w", err)
	}

	fmt.Println("--------------------")
	fmt.Println("proof of file added:")
	fmt.Println("transaction id:", add.Hash().Hex())

	return nil

}

func generateUUID(filename string) uuid.UUID {
	return uuid.NewSHA1(uuid.NameSpaceX500, []byte(filename))
}

func fileHash(file string) [32]byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err = io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	//return fmt.Printf("%x", h.Sum(nil))
	return [32]byte(h.Sum(nil))
}
