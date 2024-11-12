package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"test2/models"
	"test2/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	//llamando al metodo de encriptacion
	publicKeyPath := "keys/public_key.pem"
	publickey, err := utils.LoadPublicKey(publicKeyPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	message := "TEST"
	encryptedMessage, err := utils.EncryWithPublicKey(publickey, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypted message: %s\n", encryptedMessage)

	aesKey, err := utils.GenerateRandomAESKey(256)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("AES Key: %x\n", aesKey)

	encodedIV, err := utils.GenerateRandomIV(128)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("IV: %x\n", encodedIV)
	iv, err := base64.StdEncoding.DecodeString(encodedIV)
	if err != nil {
		fmt.Println("Error decoding IV:", err)
		return
	}
	plaintext := []byte("this is a secret message")
	ciphertext, err := utils.EncryptAES(plaintext, aesKey, iv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypted message: %x\n", ciphertext)
	// Create a new Autorization struct
	Autorization := models.Autorization{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
	encryptionKey, err := utils.Post(Autorization)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encryptionKey)
	response, err := utils.Get(encryptionKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)

}
