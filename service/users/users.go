package users

import (
	"connectly/model"
	"connectly/repository"
)

type Service interface {
	CreateOrUpdate(user model.User) (model.User, error)
	Get(user_id float64) (model.User, error)
}

type service struct {
	repository repository.Repository
}

func New(r repository.Repository) Service {
	return &service{
		r,
	}
}

func (s *service) CreateOrUpdate(user model.User) (model.User, error) {
	return s.repository.CreateOrUpdateUser(user.Username, user.FirstName, user.LastName, user.BotID, user.BotName)
}

func (s *service) Get(user_id float64) (model.User, error) {
	return s.repository.GetUser(user_id)
}
