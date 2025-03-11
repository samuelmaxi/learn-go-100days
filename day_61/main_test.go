package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByID(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestGetUserNameSuccess(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := UserService{repo: mockRepo}

	mockRepo.On("GetUserByID", 1).Return("Alice", nil)

	name, err := service.GetUsername(1)

	assert.NoError(t, err)
	assert.Equal(t, "Alice", name)
	mockRepo.AssertExpectations(t)
}

func TestGetUsernameError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := UserService{repo: mockRepo}

	mockRepo.On("GetUserByID", 1).Return("", errors.New("user not found"))

	name, err := service.GetUsername(1)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.Equal(t, "", name)
}
