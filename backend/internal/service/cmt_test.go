//go:build unit

package service

import (
	"errors"
	"testing"

	"app/test/helper"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CmtSvcSuite struct {
	suite.Suite
	ds   helper.CmtMother
	repo *helper.MockCmtRepo
	svc  CommentService
}

func TestCmtSvcSuite(t *testing.T) {
	suite.Run(t, new(CmtSvcSuite))
}

func (s *CmtSvcSuite) SetupSuite() {
	s.repo = new(helper.MockCmtRepo)
	s.svc = NewCommentService(s.repo)
	s.ds.Init()
}

func (s *CmtSvcSuite) TearDownSuite() {
}

func (s *CmtSvcSuite) TestCommentAdd() {
	// Arrange
	cmt := s.ds.Get(1)
	cmt.Id = 0
	body := CommentBody{
		PostId:   cmt.PostId,
		ParentId: cmt.ParentId,
		Content:  cmt.Content,
	}

	s.T().Run("OK", func(t *testing.T) {
		// Arrange
		cmt.PostId = 1
		body.PostId = 1
		s.repo.On("Insert", cmt).Return(nil)
		// Act
		err1 := s.svc.Add(cmt.UserId, body)
		err2 := s.svc.Add(cmt.UserId, body)
		// Assert
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		s.repo.AssertNumberOfCalls(t, "Insert", 2)
	})

	s.T().Run("PostId not found", func(t *testing.T) {
		// Arrange
		cmt.PostId = 2
		body.PostId = 2
		s.repo.On("Insert", cmt).Return(errors.New("PostId not found"))
		// Act
		err := s.svc.Add(cmt.UserId, body)
		// Assert
		assert.Error(t, err)
	})
}

func (s *CmtSvcSuite) TestCommentGet() {
	s.T().Run("Empty", func(t *testing.T) {
		// Arrange
		id := s.ds.PostEmptyId
		s.repo.On("Select", id).Return(s.ds.GetByPost(id), nil)
		// Act
		cmt, err := s.svc.GetTree(id)
		// Assert
		assert.NoError(t, err)
		assert.JSONEq(t, s.ds.PostEmpty, cmt)
	})

	s.T().Run("Single", func(t *testing.T) {
		// Arrange
		id := s.ds.PostSingleId
		s.repo.On("Select", id).Return(s.ds.GetByPost(id), nil)
		// Act
		cmt, err := s.svc.GetTree(id)
		// Assert
		assert.NoError(t, err)
		assert.JSONEq(t, s.ds.PostSingle, cmt)
	})

	s.T().Run("Nested", func(t *testing.T) {
		// Arrange
		id := s.ds.PostNestedId
		s.repo.On("Select", id).Return(s.ds.GetByPost(id), nil)
		// Act
		cmt, err := s.svc.GetTree(id)
		// Assert
		assert.NoError(t, err)
		assert.JSONEq(t, s.ds.PostNested, cmt)
	})
}
