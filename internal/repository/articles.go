package repository

import (
	"errors"

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

func (r *ArticlesRepository) Create(input domain.ArticleInput, userId int) error {
	query := "INSERT INTO articles (user_id, title, content) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, userId, input.Title, input.Content)
	return err
}

func (r *ArticlesRepository) GetById(id int) (domain.ArticleOutput, error) {
	var article domain.ArticleOutput
	query := `SELECT ar.id, CONCAT(us.name, ' ', us.surname) as author, ar.title, ar.content, ar.created_at
			  FROM articles ar INNER JOIN users us ON ar.user_id = us.id
			  WHERE ar.id = $1`
	err := r.DB.Get(&article, query, id)
	return article, err
}

func (r *ArticlesRepository) Delete(id, userId int) error {
	result, err := r.DB.Exec("DELETE FROM articles WHERE id = $1 and user_id = $2", id, userId)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("can't delete another user article")
	}
	return err
}
