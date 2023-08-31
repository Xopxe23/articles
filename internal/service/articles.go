package service

import (
	"time"

	"github.com/xopxe23/articles/internal/domain"
	"github.com/xopxe23/articles/internal/transport/rabbitmq"
)

type ArticlesRepository interface {
	GetAll() ([]domain.ArticleOutput, error)
	Create(input domain.ArticleInput, userId int) (int, error)
	GetById(id int) (domain.ArticleOutput, error)
	Update(id, userId int, input domain.ArticleInput) error
	Delete(id, userId int) error
}

type ArticlesService struct {
	articlesRepo ArticlesRepository
	rmqClient    rabbitmq.Client
}

func NewArticlesService(repo ArticlesRepository, rmq rabbitmq.Client) *ArticlesService {
	return &ArticlesService{articlesRepo: repo, rmqClient: rmq}
}

func (s *ArticlesService) GetAll() ([]domain.ArticleOutput, error) {
	return s.articlesRepo.GetAll()
}

func (s *ArticlesService) Create(input domain.ArticleInput, userId int) error {
	id, err := s.articlesRepo.Create(input, userId)
	if err != nil {
		return err
	}
	err = s.rmqClient.SendLog(rabbitmq.LogItem{
		UserId: userId,
		ArticleId: id,
		Action: "CREATE",
		Time: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ArticlesService) GetById(id int) (domain.ArticleOutput, error) {
	return s.articlesRepo.GetById(id)
}

func (s *ArticlesService) Update(id, userId int, input domain.ArticleInput) error {
	err := s.articlesRepo.Update(id, userId, input)
	if err != nil {
		return err
	}

	err = s.rmqClient.SendLog(rabbitmq.LogItem{
		UserId: userId,
		ArticleId: id,
		Action: "UPDATE",
		Time: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ArticlesService) Delete(id, userId int) error {
	err := s.articlesRepo.Delete(id, userId)
	if err != nil {
		return err
	}
	err = s.rmqClient.SendLog(rabbitmq.LogItem{
		UserId: userId,
		ArticleId: id,
		Action: "DELETE",
		Time: time.Now(),
	})
	return err
}
