-- +goose Up
CREATE TABLE users(
    id UUID,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;

