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

func (r *StartupScriptRepository) FindScriptsByProjectID(ctx context.Context, id uuid.UUID) ([]scripts.OutputScripts, error) {
	rows, err := r.db.QueryContext(ctx, `select script, path_project from startup_script where project_id = $1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	s := []scripts.OutputScripts{}

	for rows.Next() {
		var scriptAtual scripts.OutputScripts

		if err := rows.Scan(&scriptAtual.Command, &scriptAtual.Path); err != nil {
			return nil, err
		}

		s = append(s, scriptAtual)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return s, nil
}
