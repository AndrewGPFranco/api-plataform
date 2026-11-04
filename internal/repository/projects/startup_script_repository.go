package projects

import (
	"context"
	"database/sql"
	"log"
	"plataform/internal/dtos/scripts"

	"github.com/google/uuid"
)

type StartupScriptRepository struct {
	db *sql.DB
}

func (r *StartupScriptRepository) NewScript(ctx context.Context, dto scripts.InputStartupScript) error {
	query := r.db.QueryRowContext(ctx, `insert into startup_script 
    (project_id, path_project, script, created_at, updated_at) values $1, $2, $3, $4, $5`)

	return query.Scan(&dto.ProjectID, dto.PathProject, dto.Script, dto.CreatedAt, dto.UpdatedAt)
}

func NewStartupScriptRepository(db *sql.DB) *StartupScriptRepository {
	return &StartupScriptRepository{db: db}
}

func (r *StartupScriptRepository) FindCommandByProjectID(ctx context.Context, id uuid.UUID) (*string, error) {
	var command *string

	err := r.db.QueryRowContext(ctx, `select startup_script from projects where id = $1`, id).Scan(&command)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return command, nil
}
