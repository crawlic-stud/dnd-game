package router

import (
	"dnd-game/internal/models"
	"net/http"
)

func (api *Router) Login(w http.ResponseWriter, r *http.Request) error {
	var user models.LoginUser
	if err := api.GetBody(w, r, &user); err != nil {
		return err
	}

	userDb, err := api.Store.GetUserByUsername(r.Context(), user.Username)
	if err != nil {
		return api.NotFound("User '%s' not found", user.Username)
	}

	if !api.Auth.CheckPasswordHash(user.Password, userDb.HashedPassword) {
		return api.Unauthorized("Password is incorrect")
	}

	token, err := api.Auth.GenerateToken(userDb.ID.String())
	if err != nil {
		return err
	}

	return api.OK(w, map[string]any{
		"token": token,
	})
}
