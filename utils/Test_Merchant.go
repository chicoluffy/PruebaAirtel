package utils

import (
	"fmt"
	"time"
)

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

/*
func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

    url := ""
    accessToken := "your_access_token"
    username := "your_username"
    amount := "your_amount"
    transid := "your_transid"

    response, err := utils.MakeAirtelRequest(url, accessToken, username, amount, transid)
    if err != nil {
        fmt.Println("Error making Airtel request:", err)
        return
    }

    fmt.Println("Response:", response)

    serial := utils.GenerateSerial()
    fmt.Println("Generated serial:", serial)

    timestamp := utils.GetCurrentTimestamp()
    fmt.Println("Current timestamp:", timestamp)
}
*/
