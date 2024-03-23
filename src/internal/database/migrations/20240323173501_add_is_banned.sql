-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN is_banned BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN is_banned;
-- +goose StatementEnd
