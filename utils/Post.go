package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"test2/models"

	"github.com/joho/godotenv"
)

func Post(Autorization models.Autorization) (models.EncryptionKey, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return models.EncryptionKey{}, err
	}
	url := os.Getenv("URL_AUTORIZATION")
	var encryptionKey models.EncryptionKey
	jsonData, err := json.Marshal(Autorization)
	if err != nil {
		return encryptionKey, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return encryptionKey, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return encryptionKey, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return encryptionKey, err
	}

	err = json.Unmarshal(body, &encryptionKey)
	if err != nil {
		return encryptionKey, err
	}
	return encryptionKey, nil

}
