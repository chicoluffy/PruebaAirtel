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
func GetCurrentTimeStamp() int64 {
	now := time.Now()
	date := now.Format("2006-01-02")
	timeStr := now.Format("15:04:05")
	timestamp, err := time.Parse("2006-01-02 15:04:05", date+" "+timeStr)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return timestamp.Unix()
}

func MakeAirtelRequest(url, accessToken, username, amount, transid, XSignature, XKey string, requestBody map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(requestBody)
	fmt.Println("Dentro de la funcion:" + string(jsonData))
	fmt.Println("URL:" + url)
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
	req.Header.Set("X-Signature", XSignature)
	req.Header.Set("X-Key", XKey)

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
