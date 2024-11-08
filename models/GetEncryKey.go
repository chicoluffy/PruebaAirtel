package models

import (
	"time"
)

type GetEncryptionKey struct {
	Data struct {
		KeyId     int       `json:"key_id"`
		Key       string    `json:"key"`
		ValidUpTo time.Time `json:"valid_up_to"`
	} `json:"data"`
	Status struct {
		Code         string `json:"code"`
		Message      string `json:"message"`
		ResponseCode string `json:"response_code"`
		ResultCode   string `json:"result_code"`
		Success      bool   `json:"success"`
	} `json:"status"`
}
