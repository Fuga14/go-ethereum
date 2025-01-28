package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	rpcEndpoint := "https://eth-mainnet.g.alchemy.com/v2/Znfzv2rit7VpF-RL14RL1Jxafmam95a3"

	client, err := ethclient.Dial(rpcEndpoint)

	if err != nil {
		log.Fatal("Error making eth clinet", err)
	}

	fmt.Println("Connection successful")

	fmt.Println(client.BlockNumber(context.Background()))

	fmt.Println(client.BalanceAt(context.Background(), common.HexToAddress("0x7baBf95621f22bEf2DB67E500D022Ca110722FaD"), nil))
}
