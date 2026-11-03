package projects

import (
	"time"
	"uuid"
)

type Startup struct {
	ProjectID     uuid.UUID `json:"project_id"`
	PathProject   *string   `json:"path_project"`
	StartupScript *string   `json:"startup_script"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
