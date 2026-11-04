package projects

import (
	"context"
	"database/sql"
	"log"
	"time"

	"plataform/internal/dtos/scripts"

	"github.com/google/uuid"
)

type StartupScriptRepository struct {
	db *sql.DB
}

func (r *StartupScriptRepository) NewScript(ctx context.Context, dto scripts.InputStartupScript) error {
	now := time.Now()

	_, err := r.db.ExecContext(ctx, `
        insert into startup_script
            (project_id, path_project, script, title, created_at, updated_at)
        values ($1, $2, $3, $4, $5, $6)`,
		dto.ProjectID, dto.PathProject, dto.Script, dto.Title, now, now,
	)

	return err
}

func NewStartupScriptRepository(db *sql.DB) *StartupScriptRepository {
	return &StartupScriptRepository{db: db}
}

func (r *StartupScriptRepository) FindCommandByProjectID(ctx context.Context, id uuid.UUID) (command *string, path *string, err error) {
	err = r.db.QueryRowContext(ctx,
		`select script, path_project from startup_script where project_id = $1`, id,
	).Scan(&command, &path)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return command, path, nil
}
