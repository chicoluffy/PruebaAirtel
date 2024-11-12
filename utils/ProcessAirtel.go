package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"test2/models"
	"time"
)

func GenerateSerial() string {
	const max = 1679615
	rand.NewSource(time.Now().UnixNano())
	parts := make([]string, 4)
	for i := 0; i < 4; i++ {
		num := rand.Intn(max + 1)
		parts[i] = strings.ToUpper(fmt.Sprintf("%04s", strconv.FormatInt(int64(num), 36)))
	}
	return strings.Join(parts, "-")
}

func MakeAirtelRequest(url, accessToken, username, amount, transid string) (string, error) {
	requestBody := models.AirtelRequest{
		Reference: "Testingtransaction",
		Subscriber: models.Subscriber{
			Country:  "MW",
			Currency: "MWK",
			Msisdn:   username,
		},
		Transaction: models.Transaction{
			Amount:   amount,
			Country:  "MW",
			Currency: "MWK",
			ID:       transid, //cambiar poro generate serial
		},
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("X-Country", "MW")
	req.Header.Set("X-Currency", "MWK")
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
	return string(body), nil

}