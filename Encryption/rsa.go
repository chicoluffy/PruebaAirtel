package Encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func EncryWithPublicKey(publicKey *rsa.PublicKey, message []byte) (string, error) {

	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, nil)
	if err != nil {
		return "", err
	}
	encodeMessage := base64.StdEncoding.EncodeToString(ciphertext)
	return encodeMessage, nil
}

func LoadPublicKey(filepath string) (*rsa.PublicKey, error) {
	publicKeyFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer publicKeyFile.Close()

	publicKeyPEM, err := io.ReadAll(publicKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}
	return rsaPublicKey, nil
}
