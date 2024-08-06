-- +goose Up
ALTER TABLE users
ADD COLUMN apiKey VARCHAR(64) NOT NULL DEFAULT ENCODE(SHA256(random()::text::bytea),'base64');

-- +goose Down
ALTER TABLE users
DROP COLUMN apiKey;
