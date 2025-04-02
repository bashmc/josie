-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    name TEXT NOT NULL,
    size INT NOT NULL,
    ext VARCHAR(6) NULL,
    url TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
