package helper

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Detail string `json:"detail"`
}

// OK writes 200 response with json data
func (helper *ServerHelper) OK(w http.ResponseWriter, model any) error {
	return helper.HTTPResponse(w, model, http.StatusOK)
}

// NoContent writes 204 response with json data
func (helper *ServerHelper) NoContent(w http.ResponseWriter) error {
	return helper.HTTPResponse(w, nil, http.StatusNoContent)
}

// HTTPResponse writes response with model and status code
func (helper *ServerHelper) HTTPResponse(w http.ResponseWriter, model any, statusCode int) error {
	if model == nil {
		w.WriteHeader(statusCode)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(model)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	w.Write(response)
	return nil
}
