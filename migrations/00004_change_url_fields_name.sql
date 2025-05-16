-- +goose Up
-- +goose StatementBegin
ALTER TABLE urls
    RENAME COLUMN shortUrl TO short_url;

ALTER TABLE urls
    RENAME COLUMN longUrl TO long_url;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE urls
    RENAME COLUMN short_url TO shortUrl;

ALTER TABLE urls
    RENAME COLUMN long_url TO longUrl;
-- +goose StatementEnd