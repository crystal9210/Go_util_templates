package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

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
// ☆pkdf2関数は鍵導出プロセスにソルト＋イテレーションを組み込むことでより強固なセキュリティを提供→ソルト：同じPWから導出される鍵が毎回異なり、レインボーテーブル攻撃に対する耐性を高める、イテレーション：ブルートフォース攻撃に対する耐性を向上させる
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
	result := append(salt, nonce...)       // saltの末尾に追加の引数nonceの全要素を追加
	result = append(result, ciphertext...) // resultの末尾にciphertextを追加
	// ソルト＋nonce＋ciphertextを順に一つのもj利悦に連結した結果を生成し、出力している

	// 暗号化データを出力
	fmt.Printf("EncryptedData: %x\n", result)
}

func main() {
	// ユーザーAとBの鍵ペアを生成
	publicKeyA, privateKeyA, err := GenerateKeyPair()
	if err != nil {
		log.Fatalf("Error generating keys for user A: %v", err)
	}

	publicKeyB, privateKeyB, err := GenerateKeyPair()
	if err != nil {
		log.Fatalf("Error generating keys for user B: %v", err)
	}

	// 公開鍵を交換して共有秘密鍵を生成
	sharedSecretA, err := ComputeSharedSecret(privateKeyA, publicKeyB)
	if err != nil {
		log.Fatalf("Error computing shared secret for user A: %v", err)
	}

	sharedSecretB, err := ComputeSharedSecret(privateKeyB, publicKeyA)
	if err != nil {
		log.Fatalf("Error computing shared secret for user B: %v", err)
	}

	// メッセージを初期化
	message := "Hello, world!" // ここでmessage変数を初期化

	// メッセージを暗号化
	encryptedMessage, err := EncryptWithAESGCM(sharedSecretA, []byte(message))
	if err != nil {
		log.Fatalf("Error encrypting message: %v", err)
	}

	// メッセージを復号化
	decryptedMessage, err := DecryptWithAESGCM(sharedSecretB, encryptedMessage)
	if err != nil {
		log.Fatalf("Error decrypting message: %v", err)
	}

	fmt.Printf("Original Message: %s\n", message)
	fmt.Printf("Encrypted Message: %s\n", encryptedMessage)
	fmt.Printf("Decrypted Message: %s\n", decryptedMessage)
}
