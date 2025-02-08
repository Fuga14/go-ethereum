package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	rpcEndpoint := os.Getenv("RPC_ENDPOINT_ETH_SEPOLIA")

	client, err := ethclient.Dial(rpcEndpoint)

	if err != nil {
		log.Fatal("Error making eth clinet", err)
	}

	fmt.Println("Connection successful")

	fmt.Println(client.BlockNumber(context.Background()))

	fmt.Println(client.BalanceAt(context.Background(), common.HexToAddress("0x0309004C4fB9943797f5C530abd8cddE564A9fD4"), nil))
}
