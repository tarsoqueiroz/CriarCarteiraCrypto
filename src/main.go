package main

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {
	// Generate a new seed
	seed := make([]byte, 32)
	_, err := rand.Read(seed)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the master key from the seed
	masterKey, err := hdkeychain.NewMaster(seed, btcsuite.MainNet)
	if err != nil {
		log.Fatal(err)
	}

	// Print the master private key
	fmt.Println("Master Private Key:", masterKey.String())

	// Derive the first child key (m/0'/0/0)
	firstChild, err := masterKey.Child(0) // first child key
	if err != nil {
		log.Fatal(err)
	}

	// Print first child private key
	fmt.Println("First Child Private Key:", firstChild.String())

	// You can derive further child keys using the Child method
	secondChild, err := firstChild.Child(1) // derive second child key
	if err != nil {
		log.Fatal(err)
	}

	// Print second child private key
	fmt.Println("Second Child Private Key:", secondChild.String())
}
