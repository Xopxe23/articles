package service

import (
	"github.com/xopxe23/articles/internal/domain"
)

type ArticlesRepository interface {
	GetAll() ([]domain.ArticleOutput, error)
}

type ArticlesService struct {
	articlesRepo ArticlesRepository
}

func NewArticlesService(repo ArticlesRepository) *ArticlesService {
	return &ArticlesService{articlesRepo: repo}
}

func (s *ArticlesService) GetAll() ([]domain.ArticleOutput, error) {
	return s.articlesRepo.GetAll()
}
