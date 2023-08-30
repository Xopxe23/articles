package service

import (
	"github.com/xopxe23/articles/internal/domain"
)

type ArticlesRepository interface {
	GetAll() ([]domain.ArticleOutput, error)
	Create(input domain.ArticleInput, userId int) error
	GetById(id int) (domain.ArticleOutput, error)
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

func (s *ArticlesService) Create(input domain.ArticleInput, userId int) error {
	return s.articlesRepo.Create(input, userId)
}

func (s *ArticlesService) GetById(id int) (domain.ArticleOutput, error) {
	return s.articlesRepo.GetById(id)
}
