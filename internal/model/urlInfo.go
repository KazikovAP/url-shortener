package model

import "time"

type URLInfo struct {
	ID        int64
	URL       string
	Alias     string
	CreatedAt time.Time
}
