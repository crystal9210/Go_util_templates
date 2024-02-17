package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

func RandomBytes(size int) ([]byte, error) {
	dst := make([]byte, size)
	if _, err := rand.Read(dst); err != nil {
		return nil, err
	}
	return dst, nil
}

func EncryptWithAESGCM(sharedSecret, plaintext []byte) (string, error) {
	salt, _ := RandomBytes(16)
	key := pbkdf2.Key(sharedSecret, salt, 4096, 32, sha256.New)
	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	nonce, _ := RandomBytes(12)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	encryptedData := append(salt, nonce...)
	encryptedData = append(encryptedData, ciphertext...)
	return hex.EncodeToString(encryptedData), nil
}

func DecryptWithAESGCM(sharedSecret []byte, encryptedDataHex string) ([]byte, error) {
	encryptedData, err := hex.DecodeString(encryptedDataHex)
	if err != nil {
		return nil, err
	}

	if len(encryptedData) < 16+12 { // ソルト(16バイト) + ノンス(12バイト)の長さを確認
		return nil, errors.New("encrypted data too short")
	}

	salt := encryptedData[:16]
	nonce := encryptedData[16:28]
	ciphertext := encryptedData[28:]

	key := pbkdf2.Key(sharedSecret, salt, 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
