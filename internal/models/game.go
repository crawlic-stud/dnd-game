package models

import (
	"dnd-game/internal/util/validation"
	"fmt"

	"github.com/google/uuid"
)

type GameObject struct {
	ObjectID uuid.UUID `json:"objectID"`
	PosX     int16     `json:"posX"`
	PosY     int16     `json:"posY"`
}

type GameScene struct {
	Name     string       `json:"name"`
	MapImage string       `json:"mapImage"`
	Width    int16        `json:"width"`
	Height   int16        `json:"height"`
	Objects  []GameObject `json:"objects"`
}

type GameCreate struct {
	Name       string      `json:"name"`
	MaxPlayers int16       `json:"maxPlayers"`
	Scenes     []GameScene `json:"scenes"`
	HostID     uuid.UUID   `json:"hostID"`
}

func (g GameCreate) Validate() error {
	validator := validation.NewValidator(g).
		Add(g.Name != "", "name must not be empty").
		Add(g.MaxPlayers > 0, "maxPlayers must be greater than 0").
		Add(len(g.Scenes) > 0, "scenes must not be empty").
		Add(g.HostID != uuid.Nil, "hostID must not be empty")
	for i, scene := range g.Scenes {
		sceneValid := scene.Validate()
		validator = validator.Add(sceneValid == nil, fmt.Sprintf("scene %d is invalid, %v", i, sceneValid))
	}
	return validator.Validate()
}

func (s GameScene) Validate() error {
	return validation.NewValidator(s).
		Add(s.Name != "", "name must not be empty").
		Add(s.MapImage != "", "mapImage must not be empty").
		Add(s.Width > 0, "width must be greater than 0").
		Add(s.Height > 0, "height must be greater than 0").
		Validate()
}
