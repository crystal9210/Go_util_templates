package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func randomBytes(size int) []byte {
	dst := make([]byte, size)
	if _, err := rand.Read(dst); err != nil {
		panic("randomBytes error: " + err.Error())
	}
	return dst
}

// ソルトを含んで鍵を導出し、ソルトも返す
func encryptKeyAndSalt(key []byte) ([]byte, []byte) {
	salt := randomBytes(16)
	encKey := pbkdf2.Key(key, salt, 4096, 32, sha256.New)
	return encKey, salt
}

func encrypt() {
	key, _ := hex.DecodeString("4b062aa58974ab41900050afd7ad35cbb33342b10d499737ef011c44c95df218")
	encKey, salt := encryptKeyAndSalt(key)
	plaintext := []byte("hello, world!")

	block, err := aes.NewCipher(encKey)
	if err != nil {
		panic("NewCipher error: " + err.Error())
	}
	// nonce:number once;暗号化のために1度だけ使用される数をランダムに生成する関数;GCMモードではnonceは暗号化ごとに異なる必要がある
	nonce := randomBytes(12)

	// AES暗号化ブロックを使用してGCM(galois/counter mode)暗号化インスタンスを作成し、GCMは認証付き暗号モードでデータの気密性と整合性の両方を保証しておりこの下で、平文を暗号化する。
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("NewGCM error: " + err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	// ソルトとnonceを結果に追加する
	result := append(salt, nonce...)
	result = append(result, ciphertext...)

	fmt.Printf("EncryptedData: %x\n", result)
}

func main() {
	encrypt()
}
