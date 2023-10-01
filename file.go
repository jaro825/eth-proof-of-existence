package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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

type File struct {
	UUID    uuid.UUID  `json:"uuid"`
	Name    string     `json:"name"`
	HashSum string     `json:"hashsum,omitempty"`
	TxId    string     `json:"tx_id,omitempty"`
	Created *time.Time `json:"created,omitempty"`
}

func ActionFileAdd(ctx *cli.Context) error {
	privateKeyHex := getStringParamFromContext(ctx, "private_key")
	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("HexToECDSA error: %w", err)
	}

	contractAddressHex := getStringParamFromContext(ctx, "contract_address")
	contractAddress := common.HexToAddress(contractAddressHex)

	ownerAddressHex := getStringParamFromContext(ctx, "contract_owner")
	ownerAddress := common.HexToAddress(ownerAddressHex)

	filePath := getStringParamFromContext(ctx, "file")
	hash := fileHash(filePath)
	fileBasename := filepath.Base(filePath)
	fileUUID := generateUUID(fileBasename)

	nodeURL := getStringParamFromContext(ctx, "node_url")
	backend, err := NewNetworkBackend(nodeURL)
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

	timestamp := time.Unix(int64(created), 0)

	file := &File{
		UUID:    fileUUID,
		Name:    fileBasename,
		HashSum: fmt.Sprintf("%x", hash),
		TxId:    add.Hash().Hex(),
		Created: &timestamp,
	}

	fmt.Println("--------------------")
	fmt.Println("proof of file added:")
	prettyPrint(file)

	return nil
}

func getStringParamFromContext(ctx *cli.Context, param string) string {
	value := ctx.Value(param)

	if v, ok := value.(string); ok && v != "" {
		return v
	}

	panic("parameter is empty")
}

func ActionFileVerify(ctx *cli.Context) error {
	contractAddressHex := getStringParamFromContext(ctx, "contract_address")
	contractAddress := common.HexToAddress(contractAddressHex)

	filePath := getStringParamFromContext(ctx, "file")
	fileBasename := filepath.Base(filePath)
	fileUUID := generateUUID(fileBasename)
	hash := fileHash(filePath)

	verifyFile := &File{
		UUID:    fileUUID,
		Name:    fileBasename,
		HashSum: fmt.Sprintf("%x", hash),
	}

	nodeURL := getStringParamFromContext(ctx, "node_url")
	backend, err := NewNetworkBackend(nodeURL)
	if err != nil {
		return fmt.Errorf("NewNetworkBackend error: %w", err)
	}

	instance, err := NewStore(contractAddress, backend)
	if err != nil {
		return fmt.Errorf("NewStore error: %w", err)
	}

	key := [16]byte(fileUUID[:])

	fName, fHash, fCreated, err := instance.Get(&bind.CallOpts{}, key)
	if err != nil {
		return fmt.Errorf("Add error: %w", err)
	}

	name := make([]byte, len(fName))
	n, err := base64.StdEncoding.Decode(name, fName)
	if err != nil {
		return fmt.Errorf("Decode error: %w", err)
	}

	name = name[0:n]

	created := time.Unix(int64(fCreated), 0)

	checksumMatch := hash == fHash
	switch checksumMatch {
	case true:
		fmt.Println("+++ checksums match +++")
	case false:
		fmt.Println("--- checksums don't match ---")
	}

	onChainFile := &File{
		UUID:    fileUUID,
		Name:    string(name),
		HashSum: fmt.Sprintf("%x", hash),
		Created: &created,
	}

	fmt.Println("--------------------")
	fmt.Println("verify file metadata:")
	prettyPrint(verifyFile)
	fmt.Println("file metadata found on-chain:")
	prettyPrint(onChainFile)

	return nil
}

func prettyPrint(data interface{}) {
	fmt.Println("--------------------")
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	fmt.Println("--------------------")
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

	return [32]byte(h.Sum(nil))
}
