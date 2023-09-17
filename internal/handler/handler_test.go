package handler

import (
	"api/pkg/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	er "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func makeGetHandler() (*httptest.ResponseRecorder, error) {
	// make a "get" request to keys endpoint
	requestHttp, err := http.NewRequest("GET", "/api/keys", nil)
	if err != nil {
		return nil, er.Wrap(err, "error_try_get_during_test")
	}

	// create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	GetHandler(rr, requestHttp)
	return rr, nil
}

func TestGetHandlerSuccess(t *testing.T) {
	// Create a simulated HTTP request
	rr, err := makeGetHandler()
	assert.Nil(t, err)

	// Verify the HTTP status code
	assert.Equal(t, rr.Code, http.StatusOK)

	// Verify the Content-Type header
	expectedContentType := "application/json"
	contentType := rr.Header().Get("Content-Type")
	assert.Equal(t, expectedContentType, contentType)
}

func TestPostHandlerSuccess(t *testing.T) {
	rr, err := makeGetHandler()
	assert.Nil(t, err)

	// verify the HTTP status code
	assert.Equal(t, rr.Code, http.StatusOK)

	//get publickey from response
	var keys models.ResponseKeys
	err = json.Unmarshal(rr.Body.Bytes(), &keys)
	assert.Nil(t, err)
	publicKey := models.RequestDTO{PublicKey: keys.PublicKey}
	jsonData, err := json.Marshal(publicKey)
	assert.Nil(t, err)

	// create a simulated HTTP request with a valid JSON payload
	req, err := http.NewRequest("POST", "/api/address", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()

	// call the PostHandler
	PostHandler(rr, req)

	// verify the HTTP status code is OK (200) and response value
	assert.Equal(t, http.StatusOK, rr.Code)
	var address models.ResponseAddress
	err = json.Unmarshal(rr.Body.Bytes(), &address)
	assert.Nil(t, err)
	assert.NotNil(t, address)
}

func TestPostHandlerError(t *testing.T) {
	publicKeyError := models.RequestDTO{PublicKey: "error_value"}

	jsonData, err := json.Marshal(publicKeyError)
	assert.Nil(t, err)

	// create a simulated HTTP request with a invalid JSON payload
	req, err := http.NewRequest("POST", "/api/address", bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	PostHandler(rr, req)

	// verify the HTTP status code is 500 and response value is ""
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	var address models.ResponseAddress
	err = json.Unmarshal(rr.Body.Bytes(), &address)
	assert.Nil(t, err)
	assert.Equal(t, "", address.Address)
}
