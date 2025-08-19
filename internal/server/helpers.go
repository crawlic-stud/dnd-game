package server

import (
	"dnd-game/internal/util/validation"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// GetBody scans into struct and validates JSON body
func (s *Server) GetBody(w http.ResponseWriter, r *http.Request, model validation.BaseModel) error {
	err := json.NewDecoder(r.Body).Decode(model)
	defer r.Body.Close()

	if err != nil {
		var synErr *json.SyntaxError
		var unmarshalErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &synErr):
			return s.UnprocessableEntity("request body contains badly-formed JSON (at position %d)", synErr.Offset)
		case errors.Is(err, io.EOF):
			return s.UnprocessableEntity("request body must not be empty")
		case errors.As(err, &unmarshalErr):
			return s.UnprocessableEntity("request body contains an invalid value for the %q field (at position %d)", unmarshalErr.Field, unmarshalErr.Offset)
		default:
			return err
		}
	}

	if err = model.Validate(); err != nil {
		return s.UnprocessableEntity(err.Error())
	}

	return nil
}
