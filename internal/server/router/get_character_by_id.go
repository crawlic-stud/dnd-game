package router

import (
	"dnd-game/internal/util/mapper"
	"net/http"
)

func (api *Router) GetCharacterById(w http.ResponseWriter, r *http.Request) {
	characterUUID, err := api.UUIDFromPath(r, "character_id")
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	character, err := api.Store.GetCharacterByID(r.Context(), characterUUID)
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	response, err := mapper.CharacterResponse(character)
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	api.OK(w, response)
}
