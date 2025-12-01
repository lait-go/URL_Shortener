-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS analytics (
    id SERIAL PRIMARY KEY,
    short_url VARCHAR(13) NOT NULL,
    ip INET NOT NULL,
    user_agent TEXT NOT NULL,
    time TIMESTAMP NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF NOT EXISTS analytics;
-- +goose StatementEnd
