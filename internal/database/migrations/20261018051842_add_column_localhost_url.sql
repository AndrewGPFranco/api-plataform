-- +goose Up
ALTER TABLE projects
    ADD COLUMN localhost_url varchar(100) not null default 'http://localhost:8000';

-- +goose Down
ALTER TABLE projects DROP COLUMN localhost_url;