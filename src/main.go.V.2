package main

import (
	"fmt"

	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {
	// Gerar uma seed aleatória com 32 bytes
	masterSeed := make([]byte, 32)
	// ... (utilizar uma fonte de aleatoriedade segura para preencher o slice)

	// Criar a chave mestre
	masterKey, err := hdkeychain.NewMaster(masterSeed, nil)
	if err != nil {
		panic(err)
	}

	// Derivar uma chave para a nossa carteira (caminho de derivação m/44'/0'/0'/0)
	// Este caminho é comumente utilizado para carteiras HD
	childKey, err := masterKey.Child(hdkeychain.Hardened(44)).
		Child(hdkeychain.Hardened(0)).
		Child(hdkeychain.Hardened(0)).
		Child(0)
	if err != nil {
		panic(err)
	}

	// Obter o endereço público
	pubKey := childKey.PublicKey()
	address, err := pubKey.Address(btcutil.MainNet) // Especificar a rede (mainnet ou testnet)
	if err != nil {
		panic(err)
	}

	// Obter o mnemônico
	mnemonic, err := hdkeychain.MnemonicFromSeed(masterSeed)
	if err != nil {
		panic(err)
	}

	fmt.Println("Endereço público:", address)
	fmt.Println("Mnemônico:", mnemonic)
}
