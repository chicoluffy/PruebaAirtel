package main

import (
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
