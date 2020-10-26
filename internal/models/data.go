package models

type Payment struct {
	Amount     int32  `json:"amount"`
	SenderBank string `json:"sender_bank"`
}
