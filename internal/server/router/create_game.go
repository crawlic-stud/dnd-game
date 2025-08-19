package router

import (
	"dnd-game/internal/db"
	"dnd-game/internal/models"
	"dnd-game/internal/models/enums"
	"net/http"

	"github.com/google/uuid"
)

func (api *Router) CreateGame(w http.ResponseWriter, r *http.Request) error {
	var gameCreate models.GameCreate
	if err := api.GetBody(w, r, &gameCreate); err != nil {
		return err
	}

	var (
		gameID uuid.UUID
		err    error
	)

	ctx := r.Context()
	err = api.Store.Transaction(ctx, func(tx *db.Queries) error {
		// create game
		gameID, err = tx.CreateGame(ctx, db.CreateGameParams{
			Name:       gameCreate.Name,
			MaxPlayers: gameCreate.MaxPlayers,
		})
		if err != nil {
			return err
		}

		// add host user to game
		if err = tx.AddUserToGame(ctx, db.AddUserToGameParams{
			UserID: gameCreate.HostID,
			GameID: gameID,
			Role:   enums.RoleHost,
		}); err != nil {
			return err
		}

		// add scenes
		for _, scene := range gameCreate.Scenes {
			sceneID, err := tx.CreateGameScene(ctx, db.CreateGameSceneParams{
				GameID:   gameID,
				Name:     scene.Name,
				MapImage: scene.MapImage,
				Width:    scene.Width,
				Height:   scene.Height,
			})
			if err != nil {
				return err
			}

			// add objects to scenes
			for _, object := range scene.Objects {
				if err = tx.AddObjectToScene(ctx, db.AddObjectToSceneParams{
					SceneID:  sceneID,
					ObjectID: object.ObjectID,
					PosX:     object.PosX,
					PosY:     object.PosY,
				}); err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return api.OK(w, map[string]any{
		"gameID": gameID,
	})
}
