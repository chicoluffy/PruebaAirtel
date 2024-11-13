package Encryption

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomIV(size int) ([]byte, error) {
	if size%8 != 0 {
		return nil, fmt.Errorf("the size must be a multiple of 8")
	}
	iv := make([]byte, size/8)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}
