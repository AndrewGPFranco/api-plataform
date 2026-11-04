package projects

import (
	"time"

	"github.com/google/uuid"
)

type Startup struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"project_id"`
	Title       *string   `json:"title"`
	PathProject *string   `json:"path_project"`
	Script      *string   `json:"script"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
