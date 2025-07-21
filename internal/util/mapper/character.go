package mapper

import (
	"dnd-game/internal/db"
	"dnd-game/internal/models"
	"encoding/json"
)

func CharacterResponse(character db.Character) (models.CharacterResponse, error) {
	var metadata models.CharacterMetadata
	err := json.Unmarshal(character.Metadata, &metadata)
	if err != nil {
		return models.CharacterResponse{}, err
	}

	return models.CharacterResponse{
		ID:        character.ID,
		Name:      character.Name,
		Class:     character.Class,
		Level:     character.Level,
		Avatar:    character.Avatar,
		Metadata:  metadata,
		CreatedAt: character.CreatedAt.Time,
		UpdatedAt: character.UpdatedAt.Time,
	}, nil
}
