package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/xopxe23/articles/internal/domain"
)

type ArticlesRepository struct {
	DB *sqlx.DB
}

func NewArticlesRepository(db *sqlx.DB) *ArticlesRepository {
	return &ArticlesRepository{DB: db}
}

func (r *ArticlesRepository) GetAll() ([]domain.ArticleOutput, error) {
	var articles []domain.ArticleOutput
	query := `SELECT ar.id, CONCAT(us.name, ' ', us.surname) as author, ar.title, ar.content, ar.created_at
			  FROM articles ar INNER JOIN users us ON ar.user_id = us.id`
	err := r.DB.Select(&articles, query)
	return articles, err
}
