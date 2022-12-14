package apiserver

import (
	"encoding/json"
	"net/http"
	"dev11/src/model"
)

type responseOK struct {
	Message string        `json:"message"`
	Events  []model.Event `json:"events"`
}

type responseError struct {
	Error string `json:"error"`
}

func throwError(w http.ResponseWriter, status int, err error) {
	response := responseError{
		Error: err.Error(),
	}

	out, _ := json.MarshalIndent(response, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}

func writeResponse(w http.ResponseWriter, status int, msg string, events []model.Event) {
	response := responseOK{
		Message: msg,
		Events:  events,
	}

	out, _ := json.MarshalIndent(response, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}
