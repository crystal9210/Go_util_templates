package main

import "golang.org/x/crypto/curve25519"

func (k *keyPair) KeyExcahenge(publicKey []byte) ([]byte, error) {
	return curve25519.X25519(k.priKey[:], publicKey)
}
