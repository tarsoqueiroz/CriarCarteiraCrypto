package main

import (
	"fmt"
	"math/rand"

	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

func generateSeed(wordCount int) (string, error) {
	entropy := make([]byte, wordCount*11/8)
	if _, err := rand.Read(entropy); err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

func createMasterKey(mnemonic string) (*hdkeychain.ExtendedKey, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}

	master, err := hdkeychain.NewMaster(seed, nil)
	if err != nil {
		return nil, err
	}

	return master, nil
}

func deriveChildKey(parent *hdkeychain.ExtendedKey, index uint32) (*hdkeychain.ExtendedKey, error) {
	child, err := parent.Child(index)
	if err != nil {
		return nil, err
	}

	return child, nil
}

func generateKeyPair(key *hdkeychain.ExtendedKey) (string, string, error) {
	//	privateKey := key.PrivateKey.String()
	privateKey := key.P
	publicKey := key.PublicKey.String()

	return privateKey, publicKey, nil
}

func main() {
	// Gera uma seed inicial
	seed, err := generateSeed(12)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Seed:", seed)

	// Cria a chave mestre
	masterKey, err := createMasterKey(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Deriva uma nova chave mestre
	childKey, err := deriveChildKey(masterKey, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Gera um par de chaves a partir da nova chave mestre
	privateKey, publicKey, err := generateKeyPair(childKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Private key:", privateKey)
	fmt.Println("Public key:", publicKey)
}
