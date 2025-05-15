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
	query := `INSERT INTO urls (longUrl, shortUrl, createdAt) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := u.db.ExecContext(ctx, query, url.LongUrl, url.ShortUrl, url.CreatedAt)
	return err
}

func (u *UrlModel) Get(shortUrl string) (*Url, error) {
	query := `SELECT longUrl FROM urls WHERE shortUrl = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var url Url
	err := u.db.QueryRowContext(ctx, query, shortUrl).Scan(&url.LongUrl)
	if err != nil {
		return nil, err
	}
	return &url, err
}
