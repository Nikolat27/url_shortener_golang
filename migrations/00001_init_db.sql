-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls
(
    id        BIGSERIAL PRIMARY KEY,
    longUrl   VARCHAR(255),
    shortUrl  VARCHAR(150),
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd
