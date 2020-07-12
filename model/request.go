package model

type EncryptedBody struct {
	Session string `json:"session"`
	Payload string `json:"payload"`
}
