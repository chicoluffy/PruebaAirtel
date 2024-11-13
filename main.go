package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"test2/Encryption"
	"test2/models"
	"test2/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	//llamando a los metodos de encriptacion

	aesKey, err := Encryption.GenerateRandomAESKey(256)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random AES Key:", aesKey)
	iv, err := Encryption.GenerateRandomIV(128)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random IV:", iv)
	keyBase64 := base64.StdEncoding.EncodeToString(aesKey)
	ivBase64 := base64.StdEncoding.EncodeToString([]byte(iv))
	fmt.Println("AES Key Base64:", keyBase64)
	fmt.Println("IV Base64:", ivBase64)
	//mensaje a parsear en bytes
	message := map[string]interface{}{
		"reference": "1234",
		"subscriber": map[string]interface{}{
			"country":  "UG",
			"currency": "UGX",
			"msisdn":   "752604392",
		},
		"transaction": map[string]interface{}{
			"amount":   "100",
			"country":  "UG",
			"currency": "UGX",
			"id":       "test_id",
		},
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey, err := Encryption.GetPublicKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	XSignature, err := Encryption.EncryWithPublicKey(publicKey, messageBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("X-Signature:", XSignature)
	combinado := keyBase64 + ":" + ivBase64
	fmt.Println("Combinado:", combinado)
	byteToSend := []byte(combinado)
	// Encrypt the AES key and IV with the public key
	XKey, err := Encryption.EncryWithPublicKey(publicKey, byteToSend)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("X-Key:", XKey)

	// Create a new Autorization struct
	Autorization := models.Autorization{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
	encryptionKey, err := utils.Post(Autorization)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encryptionKey)
	/*//rellenar los campos para las pruebas de los metodos
	urlAirtel := os.Getenv("url_MERCHANT")
	accessToken := encryptionKey.AccessToken
	amount := "1"
	username := "992722337"
	transid := utils.GenerateSerial()

	data, err := utils.MakeAirtelRequest(urlAirtel, accessToken, username, amount, transid)
	if err != nil {
		fmt.Println("Error making Airtel request:", err)
		return
	}
	fmt.Println("Response:", data)*/

	data2, err := utils.GetBalanceEnquiry(encryptionKey)
	if err != nil {
		fmt.Println("Error making Balance Enquiry request:", err)
		return
	}
	fmt.Println("Response2:", data2)

	response, err := utils.Get(encryptionKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)

}
