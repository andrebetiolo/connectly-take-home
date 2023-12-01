package service_test

import (
	"connectly/model"
	"connectly/repository"
	"testing"

	"github.com/stretchr/testify/suite"
	"gotest.tools/assert"
)

type SuiteMock struct {
	suite.Suite
	userService *repository.RepositoryMock
}

const (
	stubUsername  = "Username"
	stubFirstName = "FirstName"
	stubLastName  = "LastName"
	stubBotID     = 9998
	stubBotName   = "BotName"
)

func TestUserServiceSuite(t *testing.T) {
	s := &SuiteMock{
		userService: repository.NewMock(),
	}

	suite.Run(t, s)
}

func (s *SuiteMock) SetupTest() {}

func (s *SuiteMock) CreateOrUpdateUser() {
	s.userService.On("CreateOrUpdateUser", stubUsername, stubFirstName, stubLastName, stubBotID, stubBotName).Return(model.User{}, nil)

	_, err := s.userService.CreateOrUpdateUser(stubUsername, stubFirstName, stubLastName, stubBotID, stubBotName)

	assert.NilError(s.T(), err)
}
