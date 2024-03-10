package models

import "time"

type Post struct {
	ID            int64
	Title         string
	Text          string
	ViewsCount    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PublishedTime time.Time
}
