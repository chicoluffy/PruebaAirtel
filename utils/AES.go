package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomAESKey(size int) ([]byte, error) {
	if size%8 != 0 {
		return nil, fmt.Errorf("the size must be a multiple of 8")
	}
	key := make([]byte, size/8)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GenerateRandomIV(size int) (string, error) {
	if size%8 != 0 {
		return "", fmt.Errorf("the size must be a multiple of 8")
	}
	iv := make([]byte, size/8)
	_, err := rand.Read(iv)
	if err != nil {
		return "", err
	}
	encodedIV := base64.StdEncoding.EncodeToString(iv)
	return encodedIV, nil
}

func EncryptAES(plaintext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	padding := block.BlockSize() - len(plaintext)%block.BlockSize()
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padtext...)

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}
