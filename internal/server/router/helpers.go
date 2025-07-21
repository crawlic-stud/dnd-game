package router

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (r *Router) UUIDFromPath(req *http.Request, pathValue string) (uuid.UUID, error) {
	gotID := req.PathValue(pathValue)
	if gotID == "" {
		return uuid.UUID{}, fmt.Errorf("missing uuid '%v' in path", pathValue)
	}

	gotUUID, err := uuid.Parse(gotID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid uuid '%v' in path", pathValue)
	}

	return gotUUID, nil
}
