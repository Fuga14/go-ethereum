package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generate private key
	privateKey, err := crypto.GenerateKey() // 64 hex characters
	if err != nil {
		log.Fatal(err)
	}

	// Make the private key in readable format
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private key: ", hexutil.Encode(privateKeyBytes)[2:])

	// Generate public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public key: ", hexutil.Encode(publicKeyBytes)[4:])

	// Address
	// Keccak-256 hash of the public key, and then we take the last 40 characters (20 bytes) and prefix it with 0x
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address: ", address)
}
