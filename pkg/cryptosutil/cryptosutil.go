package cryptosutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Encrypt returns the encrypted {data} based on the {secret} value.
func Encrypt(data []byte, secret string) ([]byte, error) {
	aesBlock, err := aes.NewCipher([]byte(MD5Hashing(secret)))
	if err != nil {
		return nil, err
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcmInstance.Seal(nonce, nonce, data, nil), nil
}

// Decrypt returns the decrypted {data} based on the {secret} value.
func Decrypt(data []byte, secret string) ([]byte, error) {
	aesBlock, err := aes.NewCipher([]byte(MD5Hashing(secret)))
	if err != nil {
		return nil, err
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return nil, err
	}

	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := data[:nonceSize], data[nonceSize:]

	decodedData, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}

// MD5Hashing returns the MD5 encoded {value}.
func MD5Hashing(value string) string {
	md5Hash := md5.Sum([]byte(value))
	return hex.EncodeToString(md5Hash[:])
}
