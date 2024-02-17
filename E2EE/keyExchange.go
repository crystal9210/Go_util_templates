package main

import "golang.org/x/crypto/curve25519"

type keyPair struct {
	priKey *[32]byte
	pubKey *[32]byte
}

func (k *keyPair) KeyExcahenge(publicKey []byte) ([]byte, error) {
	return curve25519.X25519(k.priKey[:], publicKey)
}
