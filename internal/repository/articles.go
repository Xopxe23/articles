package repository

import (
	"errors"
	"fmt"
	"strings"

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

func (r *ArticlesRepository) Create(input domain.ArticleInput, userId int) (int,error) {
	query := "INSERT INTO articles (user_id, title, content) VALUES ($1, $2, $3) RETURNING id"
	row := r.DB.QueryRow(query, userId, input.Title, input.Content)
	var id int
	err := row.Scan(&id)
	return id, err
}

func (r *ArticlesRepository) GetById(id int) (domain.ArticleOutput, error) {
	var article domain.ArticleOutput
	query := `SELECT ar.id, CONCAT(us.name, ' ', us.surname) as author, ar.title, ar.content, ar.created_at
			  FROM articles ar INNER JOIN users us ON ar.user_id = us.id
			  WHERE ar.id = $1`
	err := r.DB.Get(&article, query, id)
	return article, err
}

func (r *ArticlesRepository) Update(id, userId int, input domain.ArticleInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE articles SET %s WHERE id = $%d and user_id = $%d", setQuery, argId, argId+1)
	args = append(args, id)
	args = append(args, userId)
	result, err := r.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user don't have article with this id")
	}
	return err
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
		return errors.New("user don't have article with this id")
	}
	return err
}
