package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"test2/models"

	"github.com/joho/godotenv"
)

func Get(encryptionKey models.EncryptionKey) (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	urlTemplate := os.Getenv("URL_Encryption")

	fmt.Println("Esto es el URL:", urlTemplate)
	req, err := http.NewRequest("GET", urlTemplate, nil)
	if err != nil {
		return "", err
	}
	fmt.Printf("Access Token: %s\n", encryptionKey.AccessToken)
	req.Header.Set("Authorization", "Bearer "+encryptionKey.AccessToken)
	req.Header.Set("X-Country", "MW")
	req.Header.Set("X-Currency", "MWK")
	req.Header.Set("Authorization", "Bearer "+encryptionKey.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))
	return string(body), nil
}
