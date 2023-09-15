package handler

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"

	er "github.com/pkg/errors"
)

type ResponseDTO struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := generatePrivateAndPublicKey()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(keys)
}

func generateEthereumKeys() (*ecdsa.PrivateKey, error) {
	// choose the  elliptic curve secp256k1
	curve := elliptic.P256()

	// generate a random private key
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func generatePrivateAndPublicKey() (*ResponseDTO, error) {
	response := &ResponseDTO{}

	privateKey, err := generateEthereumKeys()
	if err != nil {
		return nil, er.Wrap(err, "error_generating_private_key")
	}

	// private key hexa format
	privateKeyHex := hex.EncodeToString(privateKey.D.Bytes())
	response.PrivateKey = privateKeyHex

	// public key hexa format
	publicKey := privateKey.PublicKey
	publicKeyBytes := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)
	response.PublicKey = publicKeyHex

	return response, nil
}
