package reviews

import (
	"connectly/model"
	"connectly/repository"
)

type Service interface {
	CreateOrUpdate(user_id, product_id, rate float64) error
	GetAll() ([]model.Review, error)
}

type service struct {
	repository repository.Repository
}

func New(r repository.Repository) Service {
	return &service{
		r,
	}
}

func (s *service) CreateOrUpdate(user_id, product_id, rate float64) error {
	return s.repository.CreateOrUpdateReview(user_id, product_id, rate)
}

func (s *service) GetAll() ([]model.Review, error) {
	return s.repository.GetAllReviews()
}
