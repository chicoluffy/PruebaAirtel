package models

type AirtelRequest struct {
	Reference   string      `json:"reference"`
	Subscriber  Subscriber  `json:"subscriber"`
	Transaction Transaction `json:"transaction"`
}
type Subscriber struct {
	Country  string `json:"country"`
	Currency string `json:"currency"`
	Msisdn   string `json:"msisdn"`
}
type Transaction struct {
	Amount   string `json:"amount"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
	ID       string `json:"id"`
}
