package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"test2/models"
)

func Post(Autorization models.Autorization) (models.EncryptionKey, error) {
	var encryptionKey models.EncryptionKey
	jsonData, err := json.Marshal(Autorization)
	if err != nil {
		return encryptionKey, err
	}

	req, err := http.NewRequest("POST", "https://openapiuat.airtel.africa/auth/oauth2/token", bytes.NewBuffer(jsonData))
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
