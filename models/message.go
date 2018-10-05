package models

// Message model
type Message struct {
	Message string `json:"message"`
	Code    uint8  `json:"code"`
}
