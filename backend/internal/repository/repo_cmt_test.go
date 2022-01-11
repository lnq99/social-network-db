//go:build unit

package repository

import (
	"database/sql"
	"errors"
	"testing"

	"app/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CmtRepoSuite struct {
	suite.Suite
	db      *sql.DB
	dbmock  sqlmock.Sqlmock
	cmtRepo CommentRepo
}

func TestCmtRepoSuite(t *testing.T) {
	suite.Run(t, new(CmtRepoSuite))
}

func (s *CmtRepoSuite) SetupSuite() {
	var err error
	s.db, s.dbmock, err = sqlmock.New()
	if err != nil {
		s.Error(err)
	}
	s.cmtRepo = NewCommentRepo(s.db)
}

func (s *CmtRepoSuite) TearDownSuite() {
	s.db.Close()
}

func (s *CmtRepoSuite) TestSelectCmt() {
	cmtColumns := []string{"id", "userid", "postid", "parentid", "content", "created"}

	s.T().Run("OK", func(t *testing.T) {
		// Arrange
		cmtMockRows := sqlmock.NewRows(cmtColumns).AddRow(1, 2, 3, 0, "comment 1", "")
		s.dbmock.ExpectQuery("select \\* from Comment where postId=\\$1").
			WithArgs(3).WillReturnRows(cmtMockRows)

		// Act
		_, err1 := s.cmtRepo.Select(3)
		err2 := s.dbmock.ExpectationsWereMet()

		// Assert
		assert.NoError(t, err1)
		assert.NoError(t, err2)
	})

	s.T().Run("Empty", func(t *testing.T) {
		// Arrange
		cmtEmptyMockRow := sqlmock.NewRows(cmtColumns)
		s.dbmock.ExpectQuery("select \\* from Comment where postId=\\$1").
			WithArgs(2).WillReturnRows(cmtEmptyMockRow)

		// Act
		_, err1 := s.cmtRepo.Select(2)
		err2 := s.dbmock.ExpectationsWereMet()

		// Assert
		assert.NoError(t, err1)
		assert.NoError(t, err2)
	})
}

func (s *CmtRepoSuite) TestInsertCmt() {
	cmt := model.Comment{
		UserId:   1,
		PostId:   2,
		ParentId: 0,
		Content:  "comment 1",
	}

	s.T().Run("OK", func(t *testing.T) {
		// Arrange
		s.dbmock.ExpectExec("insert into Comment\\(userId, postId, parentId, content\\) values \\(\\$1, \\$2, \\$3, \\$4\\)").
			WithArgs(cmt.UserId, cmt.PostId, cmt.ParentId, cmt.Content).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.dbmock.ExpectExec("update Post set cmtCount=cmtCount\\+1 where id=\\$1").
			WithArgs(cmt.PostId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Act
		err1 := s.cmtRepo.Insert(cmt)
		err2 := s.dbmock.ExpectationsWereMet()

		// Assert
		assert.NoError(t, err1)
		assert.NoError(t, err2)
	})

	s.T().Run("Can't insert", func(t *testing.T) {
		// Arrange
		s.dbmock.ExpectExec("insert into Comment\\(userId, postId, parentId, content\\) values \\(\\$1, \\$2, \\$3, \\$4\\)").
			WithArgs(cmt.UserId, cmt.PostId, cmt.ParentId, cmt.Content).
			WillReturnError(errors.New(""))

		// Act
		err1 := s.cmtRepo.Insert(cmt)
		err2 := s.dbmock.ExpectationsWereMet()

		// Assert
		assert.Error(t, err1)
		assert.NoError(t, err2)
	})
}
