package router

import (
	"dnd-game/internal/db"
	"dnd-game/internal/models"
	"net/http"
)

func (api *Router) Register(w http.ResponseWriter, r *http.Request) error {
	var user models.LoginUser
	if err := api.GetBody(w, r, &user); err != nil {
		return err
	}

	exists, err := api.Store.UsernameExists(r.Context(), user.Username)
	if err != nil {
		return err
	}

	if exists {
		return api.Conflict("Username '%s' already exists", user.Username)
	}

	hashedPassword, err := api.Auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if err = api.Store.CreateUser(r.Context(), db.CreateUserParams{
		Username:       user.Username,
		HashedPassword: hashedPassword,
	}); err != nil {
		return err
	}

	return api.OK(w, user)
}
