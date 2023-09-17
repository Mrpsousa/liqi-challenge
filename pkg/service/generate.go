package service

import (
	m "api/pkg/models"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	er "github.com/pkg/errors"
)

var Cache = make(map[string]*ecdsa.PublicKey)

func GenerateECDSAKeys() (*m.ResponseKeys, error) {
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
	return &m.ResponseKeys{PrivateKey: hexutil.Encode(privateKeyBytes), PublicKey: pbkey}, nil
}

func GenerateAddress(pbkey string) (*m.ResponseAddress, error) {
	// verify if exists a ecdsa Public Key by pbkey string
	ecdsaPublicKey, exists := Cache[pbkey]

	if !exists {
		err := er.New("can't find public key, plz create a new one")
		customEr := er.Wrap(err, "INFO: unable_to_find_key")
		log.Println(customEr)
		return nil, err
	}

	// generate te address based on public key given
	return &m.ResponseAddress{Address: crypto.PubkeyToAddress(*ecdsaPublicKey).Hex()}, nil
}
