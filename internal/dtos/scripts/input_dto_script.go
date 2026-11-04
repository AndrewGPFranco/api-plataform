package scripts

import (
	"time"

	"github.com/google/uuid"
)

type InputStartupScript struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"project_id"`
	PathProject *string   `json:"path_project"`
	Script      *string   `json:"script"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
