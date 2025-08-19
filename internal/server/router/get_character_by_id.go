package router

import (
	"dnd-game/internal/util/mapper"
	"net/http"
)

func (api *Router) GetCharacterById(w http.ResponseWriter, r *http.Request) error {
	characterUUID, err := api.UUIDFromPath(r, "character_id")
	if err != nil {
		return api.BadRequest(err.Error())
	}

	character, err := api.Store.GetCharacterByID(r.Context(), characterUUID)
	if err != nil {
		return err
	}

	response, err := mapper.CharacterResponse(character)
	if err != nil {
		return err
	}

	return api.OK(w, response)
}
