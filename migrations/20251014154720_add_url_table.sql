-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS url (
    id SERIAL PRIMARY KEY,
    full_url TEXT UNIQUE NOT NULL,
    short_url VARCHAR(13) UNIQUE NOT NULL 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS url;
-- +goose StatementEnd
