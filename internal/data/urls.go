package data

import (
	"context"
	"database/sql"
	"time"
)

type UrlModel struct {
	db *sql.DB
}

type Url struct {
	LongUrl   string
	ShortUrl  string
	CreatedAt time.Time
}

func (u *UrlModel) Insert(url *Url) error {
	query := `INSERT INTO urls (long_url, short_url) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := u.db.ExecContext(ctx, query, url.LongUrl, url.ShortUrl)
	return err
}

func (u *UrlModel) Get(shortUrl string) (*Url, error) {
	query := `SELECT long_url FROM urls WHERE short_url = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var url Url
	err := u.db.QueryRowContext(ctx, query, shortUrl).Scan(&url.LongUrl)
	if err != nil {
		return nil, err
	}
	return &url, err
}
