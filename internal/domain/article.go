package domain

import "time"

type Article struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleInput struct {
	Title     *string    `json:"title" binding:"required,gte=10"`
	Content   *string    `json:"content" binding:"required,gte=20"`
}

type ArticleOutput struct {
	Id        int       `db:"id"`
	Author    string    `db:"author"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
