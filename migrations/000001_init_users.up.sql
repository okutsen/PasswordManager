CREATE TABLE IF NOT EXISTS users
(
    id uuid primary key,
    name text not null,
    email text unique not null,
    login text unique not null,
    password text not null,
    phone text unique,
    created_at timestamp,
    updated_at timestamp
);
