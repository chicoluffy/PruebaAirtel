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
	urlkey := os.Getenv("URL_Encryption")
	url := urlkey
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+encryptionKey.AccessToken)
	req.Header.Set("X-Country", "UG")
	req.Header.Set("X-Currency", "UGX")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
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
