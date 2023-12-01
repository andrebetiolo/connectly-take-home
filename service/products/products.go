package products

import (
	"connectly/model"
	"connectly/repository"
)

type Service interface {
	Get(product_id float64) (model.Product, error)
}

type service struct {
	repository repository.Repository
}

func New(r repository.Repository) Service {
	return &service{
		r,
	}
}

func (s *service) Get(product_id float64) (model.Product, error) {
	return s.repository.GetProduct(product_id)
}
