CREATE TABLE IF NOT EXISTS records
(
    id text primary key,
    name text unique not null,
    login text unique not null,
    password text not null,
    url text,
    description text,
    updated_by text not null,
    created_by text not null,
    updated_at timestamp not null,
    created_at timestamp not null
);
