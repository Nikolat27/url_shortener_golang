-- +goose Up
-- +goose StatementBegin
ALTER TABLE urls
RENAME COLUMN createdAt TO created_at
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
