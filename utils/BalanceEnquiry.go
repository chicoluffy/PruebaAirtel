package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"test2/models"

	"github.com/joho/godotenv"
)

func GetBalanceEnquiry(encryptionKey models.EncryptionKey) (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", fmt.Errorf("rror loading .env file")
	}
	accessToken := encryptionKey.AccessToken
	// Get the URL from the environment
	url := os.Getenv("URL_ACCOUNT")
	if url == "" {
		return "", fmt.Errorf("URL_ACCOUNT is not set")
	}

	//solicitud http get
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request")
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Country", "UG")
	req.Header.Set("X-Currency", "UGX")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response")
	}
	return string(body), nil

}
