package router

import (
	"dnd-game/internal/db"
	"dnd-game/internal/models"
	"dnd-game/internal/util/mapper"
	"encoding/json"
	"net/http"
)

func (api *Router) UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	characterUUID, err := api.UUIDFromPath(r, "character_id")
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	var updateCharacter models.CharacterCreate
	if ok := api.GetBody(w, r, &updateCharacter); !ok {
		return
	}

	metadata, err := json.Marshal(updateCharacter.Metadata)
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	updatedCharacter, err := api.Store.UpdateCharacter(r.Context(), db.UpdateCharacterParams{
		ID:       characterUUID,
		Name:     updateCharacter.Name,
		Class:    updateCharacter.Class,
		Level:    updateCharacter.Level,
		Avatar:   updateCharacter.Avatar,
		Metadata: metadata,
	})
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	response, err := mapper.CharacterResponse(updatedCharacter)
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	api.OK(w, response)
}
