package handler

import (
	m "api/pkg/models"
	svc "api/pkg/service"

	"encoding/json"
	"log"
	"net/http"

	er "github.com/pkg/errors"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := svc.GenerateECDSAKeys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(keys)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	data := &m.RequestDTO{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	post, err := svc.GenerateAddress(data.PublicKey)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		// creates and encode json error
		errorMessage := m.ErrorResponse{Error: err.Error()}
		if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
			log.Println(er.Wrap(err, "INFO: error_enconding_json_with_error"))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
