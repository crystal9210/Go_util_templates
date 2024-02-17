package main

import (
	"crypto/rand"
	"log"

	"golang.org/x/crypto/curve25519"
)

type keyPair struct {
	priKey *[32]byte
	pubKey *[32]byte
}

// 鍵生成機能のエラーハンドリングの改善
func generateKey() *keyPair {
	privateKey, publicKey := new([32]byte), new([32]byte)
	if _, err := rand.Read(privateKey[:]); err != nil {
		log.Fatalf("generateKey error: %v", err)
	}
	curve25519.ScalarBaseMult(publicKey, privateKey)
	return &keyPair{
		priKey: privateKey,
		pubKey: publicKey,
	}
}
