package model

import "time"

type Comment struct {
	ID        int64     `db:"id" json:"id"`
	UserID    *int64    `db:"user_id" json:"user_id"`
	ArticleID *int64    `db:"article_id" json:"article_id"`
	Body      string    `db:"body" json:"body"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
