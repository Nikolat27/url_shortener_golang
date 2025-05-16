-- +goose Up
-- +goose StatementBegin
ALTER TABLE urls
DROP
COLUMN expires_at,
       DROP
COLUMN max_usage,
              DROP
COLUMN password;
                     
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
