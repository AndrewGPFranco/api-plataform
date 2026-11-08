-- +goose Up
alter table projects add column user_id INTEGER;

-- +goose Down
alter table projects drop column user_id;
