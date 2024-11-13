package Encryption

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// GetPublicKey is a function that returns the public key
func GetPublicKey() (*rsa.PublicKey, error) {
	publicKeyFile, err := os.ReadFile("keys/public_key.pem")
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(publicKeyFile)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	return rsaPublicKey, nil
}
