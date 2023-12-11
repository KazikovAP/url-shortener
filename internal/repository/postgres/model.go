package postgres

import (
	"time"

	"url-shortener/internal/model"
)

type urlInfo struct {
	ID        int64     `db:"id"`
	URL       string    `db:"url"`
	Alias     string    `db:"alias"`
	CreatedAt time.Time `db:"created_at"`
}

func (u *urlInfo) ToModoel() model.URLInfo {
	return model.URLInfo{
		ID:        u.ID,
		URL:       u.URL,
		Alias:     u.Alias,
		CreatedAt: u.CreatedAt,
	}
}
