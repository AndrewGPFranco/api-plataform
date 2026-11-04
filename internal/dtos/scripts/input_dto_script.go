package scripts

import (
	"github.com/google/uuid"
)

type InputStartupScript struct {
	ProjectID   uuid.UUID `json:"project_id"`
	PathProject *string   `json:"path_project"`
	Script      *string   `json:"script"`
	Title       *string   `json:"title"`
}
