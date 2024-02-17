package main

import (
	"crypto/rand"

	"golang.org/x/crypto/curve25519"
)

func GenerateKeyPair() ([]byte, []byte, error) {
	var publicKey, privateKey [32]byte
	if _, err := rand.Read(privateKey[:]); err != nil {
		return nil, nil, err
	}
	curve25519.ScalarBaseMult(&publicKey, &privateKey)
	return publicKey[:], privateKey[:], nil
}

func ComputeSharedSecret(privateKey, publicKey []byte) ([]byte, error) {
	var priv [32]byte
	var pub [32]byte
	copy(priv[:], privateKey)
	copy(pub[:], publicKey)
	var sharedSecret [32]byte
	curve25519.ScalarMult(&sharedSecret, &priv, &pub)
	return sharedSecret[:], nil
}
