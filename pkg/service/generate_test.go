package service

import (
	m "api/internal/models"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeysSuccess(t *testing.T) {
	keyReponse, err := GenerateECDSAKeys()
	assert.Nil(t, err)
	assert.NotNil(t, keyReponse)
	assert.NotNil(t, keyReponse.PrivateKey)
	assert.NotNil(t, keyReponse.PublicKey)

	// Verify the type
	assert.Equal(t, reflect.TypeOf(keyReponse), reflect.TypeOf(&m.ResponseKeys{}))
}

func TestGenerateAddressSuccess(t *testing.T) {
	keyReponse, err := GenerateECDSAKeys()
	assert.Nil(t, err)
	assert.NotNil(t, keyReponse.PublicKey)

	// Attempt to generate the address
	responseAddress, err := GenerateAddress(keyReponse.PublicKey)
	assert.Nil(t, err)
	assert.NotNil(t, responseAddress)

}

func TestGenerateAddressError(t *testing.T) {
	wrongValue := "test"
	expectedError := "can't find public key, plz create a new one"

	responseAddress, err := GenerateAddress(wrongValue)
	assert.NotNil(t, err)
	assert.Nil(t, responseAddress)
	assert.Equal(t, expectedError, err.Error())

}
