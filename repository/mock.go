package repository

import (
	"connectly/model"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func NewMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (m *RepositoryMock) CreateOrUpdateUser(username, first_name, last_name string, bot_id int64, bot_name string) (model.User, error) {
	args := m.Called(username, first_name, last_name, bot_id, bot_name)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *RepositoryMock) GetUser(user_id string) (model.User, error) {
	args := m.Called(user_id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *RepositoryMock) GetProduct(product_id string) (model.Product, error) {
	args := m.Called(product_id)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *RepositoryMock) CreateOrUpdateReview(user_id, product_id, rate int64) error {
	args := m.Called(user_id, product_id, rate)
	return args.Error(0)
}

func (m *RepositoryMock) GetAllReviews() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}
