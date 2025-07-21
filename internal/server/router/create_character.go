package router

import (
	"dnd-game/internal/db"
	"dnd-game/internal/models"
	"dnd-game/internal/util/mapper"
	"encoding/json"
	"net/http"
)

func (api *Router) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var characterCreate models.CharacterCreate
	if ok := api.GetBody(w, r, &characterCreate); !ok {
		return
	}

	characterMetadata, err := json.Marshal(characterCreate.Metadata)
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	userID, err := api.Auth.GetUserID(r)
	if err != nil {
		api.InternalServerError(w, err)
		return
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
