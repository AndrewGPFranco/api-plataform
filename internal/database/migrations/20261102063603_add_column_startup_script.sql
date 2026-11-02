-- +goose Up
alter table projects add column startup_script varchar(255);

-- +goose Down
alter table projects drop column startup_script;
