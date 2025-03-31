-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files (

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
