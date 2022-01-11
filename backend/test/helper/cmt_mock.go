package helper

import (
	"app/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockCmtRepo struct {
	mock.Mock
}

func (m *MockCmtRepo) Select(postId int) ([]model.Comment, error) {
	args := m.Called(postId)
	return args.Get(0).([]model.Comment), args.Error(1)
}

func (m *MockCmtRepo) Insert(cmt model.Comment) error {
	args := m.Called(cmt)
	return args.Error(0)
}
