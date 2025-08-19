package router

import (
	"dnd-game/internal/db"
	"dnd-game/internal/models"
	"dnd-game/internal/util/mapper"
	"encoding/json"
	"net/http"
)

func (api *Router) UpdateCharacter(w http.ResponseWriter, r *http.Request) error {
	characterUUID, err := api.UUIDFromPath(r, "character_id")
	if err != nil {
		return api.BadRequest(err.Error())
	}

	var updateCharacter models.CharacterCreate
	if err := api.GetBody(w, r, &updateCharacter); err != nil {
		return err
	}

	metadata, err := json.Marshal(updateCharacter.Metadata)
	if err != nil {
		return err
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
		return err
	}

	response, err := mapper.CharacterResponse(updatedCharacter)
	if err != nil {
		return err
	}

	return api.OK(w, response)
}
