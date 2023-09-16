package handler

import (
	"crypto/ecdsa"
	"fmt"

	"encoding/json"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	er "github.com/pkg/errors"
)

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

var Cache = make(map[string]*ecdsa.PublicKey)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := generateECDSAKeys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(keys)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	data := &RequestDTO{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	post, err := generateAddress(data.PublicKey)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		// creates and encode json error
		errorMessage := ErrorResponse{Error: err.Error()}
		if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
			log.Println(er.Wrap(err, "INFO: error_enconding_json_with_error"))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func generateECDSAKeys() (*ResponseKeys, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Println(er.Wrap(err, "INFO: generate_pvkey_error"))
		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err := er.New("publicKey is not of type *ecdsa.publickey")
		log.Println(er.Wrap(err, "INFO: assert_error"))
		return nil, err
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key:", hexutil.Encode(publicKeyBytes))

	//generate the ecdsa public key
	pbkey := hexutil.Encode(publicKeyBytes)
	Cache[pbkey] = publicKeyECDSA
	return &ResponseKeys{PrivateKey: hexutil.Encode(privateKeyBytes), PublicKey: pbkey}, nil
}

func generateAddress(pbkey string) (string, error) {
	// verify if exists a ecdsa Public Key by pbkey string
	ecdsaPublicKey, exists := Cache[pbkey]

	if !exists {
		err := er.New("can't find public key, plz create a new one")
		customEr := er.Wrap(err, "INFO: unable_to_find_key")
		log.Println(customEr)
		return "", err
	}

	// generate te address based on public key given
	return crypto.PubkeyToAddress(*ecdsaPublicKey).Hex(), nil
}
