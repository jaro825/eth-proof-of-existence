package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type NetworkBackend struct {
	*ethclient.Client
	ChainId *big.Int
}

func NewNetworkBackend(nodeURL string) (*NetworkBackend, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &NetworkBackend{
		Client:  client,
		ChainId: chainID,
	}, nil
}
