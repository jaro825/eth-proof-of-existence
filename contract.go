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

func ActionContractDeploy(ctx *cli.Context) error {
	privateKeyHex := ctx.Value("private_key")

	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex.(string))
	if err != nil {
		return fmt.Errorf("HexToECDSA error: %w", err)
	}

	walletAddressHex := ctx.Value("address")

	address := common.HexToAddress(walletAddressHex.(string))

	nodeURL := ctx.Value("node_url")

	backend, err := NewNetworkBackend(nodeURL.(string))
	if err != nil {
		return fmt.Errorf("NewNetworkBackend error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()

	nonce, err := backend.PendingNonceAt(reqCtx, address)
	if err != nil {
		return fmt.Errorf("PendingNonceAt error: %w", err)
	}

	gasPrice, err := backend.SuggestGasPrice(context.Background())
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

	fmt.Println("--------------------")
	fmt.Println("contract deployed:")
	fmt.Println("contract address:", contractAddress)
	fmt.Println("transaction id:", tx.Hash().Hex())

	return nil
}

func ActionContractStatus(ctx *cli.Context) error {
	contractAddressHex := ctx.Value("address")

	address := common.HexToAddress(contractAddressHex.(string))

	privateKeyHex := ctx.Value("private_key")

	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex.(string))
	if err != nil {
		return fmt.Errorf("HexToECDSA error: %w", err)
	}

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

	nonce, err := backend.PendingNonceAt(reqCtx, address)
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

	fmt.Println("--------------------")
	fmt.Println("contract status:")
	fmt.Println("owner address:", ownerAddress.Hex())
	fmt.Println("number of media files proofs stored:", size)

	return nil
}
