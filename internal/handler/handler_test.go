package handler

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockRandReader is an implementation of rand.Reader that always returns an error.
type MockRandReader struct{}

func (r MockRandReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("expected_error")
}

func TestGenerateEthereumKeysSuccess(t *testing.T) {
	privateKey, err := generateEthereumKeys()
	assert.Nil(t, err)
	assert.NotNil(t, privateKey)
	// Verify the type
	assert.Equal(t, reflect.TypeOf(privateKey), reflect.TypeOf(&ecdsa.PrivateKey{}))
}

func TestGenerateEthereumKeysError(t *testing.T) {
	originalRandReader := rand.Reader

	// Replace rand.Reader with our simulated reader
	rand.Reader = MockRandReader{}

	// Defer the restoration of rand.Reader
	defer func() {
		rand.Reader = originalRandReader
	}()

	// Attempt to generate the private key
	privateKey, err := generateEthereumKeys()
	assert.NotNil(t, err)
	assert.Nil(t, privateKey)
	assert.Equal(t, "expected_error", err.Error())

}

func TestGeneratePrivateAndPublicKeySuccess(t *testing.T) {
	keys, err := generatePrivateAndPublicKey()
	assert.Nil(t, err)
	assert.NotNil(t, keys.PrivateKey)
	assert.NotNil(t, keys.PublicKey)
}

func TestGeneratePrivateAndPublicKeyError(t *testing.T) {
	originalRandReader := rand.Reader

	// Replace rand.Reader with our simulated reader
	rand.Reader = MockRandReader{}

	// Defer the restoration of rand.Reader
	defer func() {
		rand.Reader = originalRandReader
	}()

	result, err := generatePrivateAndPublicKey()
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "error_generating_private_key: expected_error", err.Error())
}
