package models

type ResponseKeys struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type RequestDTO struct {
	PublicKey string `json:"public_key"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ResponseAddress struct {
	Address string `json:"address"`
}
