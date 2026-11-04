package projects

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type StartupScriptRepository struct {
	db *sql.DB
}

func NewStartupScriptRepository(db *sql.DB) *StartupScriptRepository {
	return &StartupScriptRepository{db: db}
}

func (r *ProjectRepository) FindCommandByProjectID(ctx context.Context, id uuid.UUID) (*string, error) {
	var command *string

	err := r.db.QueryRowContext(ctx, `select startup_script from projects where id = $1`, id).Scan(&command)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return command, nil
}
