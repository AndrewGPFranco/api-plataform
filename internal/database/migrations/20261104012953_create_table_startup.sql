-- +goose Up
create table startup_script
(
    id             uuid        not null primary key default gen_random_uuid(),
    project_id     uuid        not null references projects (id) on delete cascade,
    path_project   text,
    script         text,
    title text,
    created_at     timestamptz not null default now(),
    updated_at     timestamptz not null default now()
);

-- +goose Down
drop table startup_script;
