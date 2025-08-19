package router

import (
	"net/http"
)

func (api *Router) DeleteCharacter(w http.ResponseWriter, r *http.Request) error {
	characterUUID, err := api.UUIDFromPath(r, "character_id")
	if err != nil {
		return api.BadRequest(err.Error())
	}

	err = api.Store.DeleteCharacter(r.Context(), characterUUID)
	if err != nil {
		return err
	}

	return api.NoContent(w)
}
