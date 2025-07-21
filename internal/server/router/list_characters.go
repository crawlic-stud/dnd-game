package router

import (
	"net/http"

	"github.com/google/uuid"
)

func (api *Router) ListCharacters(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	var userUUID uuid.UUID
	var err error

	if userID != "" {
		userUUID, err = uuid.Parse(userID)
		if err != nil {
			api.ValidationError(w, err)
			return
		}
	} else {
		userUUID, err = api.Auth.GetUserID(r)
		if err != nil {
			api.InternalServerError(w, err)
			return
		}
	}

	characters, err := api.Store.GetCharactersByUserID(r.Context(), userUUID)
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	api.OK(w, characters)
}
