-- +goose Up
-- +goose StatementBegin
ALTER TABLE urls
    ADD COLUMN password VARCHAR(255),
ADD COLUMN max_usage INT,
ADD COLUMN expires_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() + INTERVAL '7 days';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE urls
DROP COLUMN IF EXISTS password,
DROP COLUMN IF EXISTS max_usage,
DROP COLUMN IF EXISTS expires_at;
-- +goose StatementEnd
