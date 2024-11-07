package main

import (
	"fmt"
	"test2/models"
	"test2/utils"
)

func main() {

	// Create a new Autorization struct
	Autorization := models.Autorization{
		ClientID:     "2a779896-3bf3-4753-a453-3f80fa85faf7",
		ClientSecret: "a4e43cc1-94a6-43eb-b749-0f53975bef26",
		GrantType:    "client_credentials",
	}
	encryptionKey, err := utils.Post(Autorization)
	if err != nil {
		fmt.Println(err)
		return
	}

}
