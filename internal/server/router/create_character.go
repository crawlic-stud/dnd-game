package router

import (
	"dnd-game/internal/db"
	"dnd-game/internal/models"
	"dnd-game/internal/util/mapper"
	"encoding/json"
	"net/http"
)

func (api *Router) CreateCharacter(w http.ResponseWriter, r *http.Request) error {
	var characterCreate models.CharacterCreate
	if err := api.GetBody(w, r, &characterCreate); err != nil {
		return err
	}

	characterMetadata, err := json.Marshal(characterCreate.Metadata)
	if err != nil {
		return err
	}

	userID, err := api.Auth.GetUserID(r)
	if err != nil {
		return err
	}

	character, err := api.Store.CreateCharacter(r.Context(), db.CreateCharacterParams{
		Name:     characterCreate.Name,
		Class:    characterCreate.Class,
		Level:    characterCreate.Level,
		Avatar:   characterCreate.Avatar,
		Metadata: characterMetadata,
		UserID:   userID,
	})

	if err != nil {
		return err
	}

	response, err := mapper.CharacterResponse(character)
	if err != nil {
		return err
	}

	return api.OK(w, response)
}
