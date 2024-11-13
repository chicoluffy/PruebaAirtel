package Encryption

import (
	"crypto/rand"
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
