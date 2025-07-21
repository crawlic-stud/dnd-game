package router

import (
	"net/http"
)

func (api *Router) DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	characterUUID, err := api.UUIDFromPath(r, "character_id")
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	err = api.Store.DeleteCharacter(r.Context(), characterUUID)
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	api.NoContent(w)
}
