package main

import (
	"fmt"

	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {
	// Mnemônico (substitua por um mnemônico válido)
	mnemonic := "your secure mnemonic phrase"

	// Criando a seed a partir do mnemônico
	seed, err := hdkeychain.NewSeed(mnemonic, "")
	if err != nil {
		panic(err)
	}

	// Criando a chave principal a partir da seed
	master, err := hdkeychain.NewMaster(seed, nil)
	if err != nil {
		panic(err)
	}

	// Derivando um caminho (exemplo: m/44'/0'/0'/0)
	// Consulte a hierarquia BIP44 para mais detalhes
	path := "m/44'/0'/0'/0"
	key, err := master.Child(path)
	if err != nil {
		panic(err)
	}

	// Obtendo o endereço público
	pubKey := key.PublicKey()
	address, err := pubKey.Address()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Endereço: %s\n", address)
}
