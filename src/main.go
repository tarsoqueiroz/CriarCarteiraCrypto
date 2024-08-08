package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {
	// Master seed aleatório
	masterSeed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		panic(err)
	}
	fmt.Println("Master seed:", masterSeed)

	// Criando uma chave principal a partir da seed
	// masterKey, err := hdkeychain.NewMaster(masterSeed, nil)
	masterKey, err := hdkeychain.NewMaster(masterSeed, &chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}

	// Derivando uma chave para a nossa carteira
	childKey, err := masterKey.Child(0)
	if err != nil {
		panic(err)
	}

	fmt.Println("childKey:", childKey)

	// Obtendo o endereço público
	// pubKey := childKey.PublicKey()
	pubKey, err := childKey.ECPubKey()
	// pubKey, err := childKey.ECPubKey()
	// address, err := pubKey.Address()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Endereço público:", address)
	fmt.Println("Endereço público2:", pubKey)
	fmt.Println("Mnemônico:", masterSeed)
}
