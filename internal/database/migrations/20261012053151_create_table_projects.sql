-- +goose Up
create table projects
(
    id               uuid        not null primary key default gen_random_uuid(),
    title            text        not null,
    description      text,
    started_at       date        not null,
    repo_front_url   text,
    repo_back_url    text,
    deploy_front_url text,
    deploy_back_url  text,
    has_deploy       boolean     not null generated always as (
        deploy_front_url is not null or deploy_back_url is not null
        ) stored,
    created_at       timestamptz not null             default now(),
    updated_at       timestamptz not null             default now()
);

-- +goose Down
drop table projects;
