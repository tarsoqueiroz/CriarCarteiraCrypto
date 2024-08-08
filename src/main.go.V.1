package main

import (
	"fmt"

	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {
	// Master seed aleatório
	masterSeed, err := hdkeychain.GenerateSeed(32)
	if err != nil {
		panic(err)
	}

	// Criando uma chave principal a partir da seed
	masterKey, err := hdkeychain.NewMaster(masterSeed, nil)
	if err != nil {
		panic(err)
	}

	// Derivando uma chave para a nossa carteira
	childKey, err := masterKey.Child(0)
	if err != nil {
		panic(err)
	}

	// Obtendo o endereço público
	// pubKey := childKey.PublicKey()
	pubkey, err := childKey.ECPubKey()
	// address, err := pubKey.Address()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Endereço público:", address)
	fmt.Println("Endereço público2:", pubkey)
	fmt.Println("Mnemônico:", masterSeed)
}
