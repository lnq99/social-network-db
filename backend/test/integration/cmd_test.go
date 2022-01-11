//go:build integration

package integration

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"app/config"
	v1 "app/internal/controller/v1"
	"app/internal/repository"
	"app/internal/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// DBMock - Repo - Service - Controller
type CommentSuite struct {
	suite.Suite
	db     *sql.DB
	dbmock sqlmock.Sqlmock
	repo   *repository.Repo
	svc    *service.Services
	ctrl   *v1.Controller
}

func TestCommentSuite(t *testing.T) {
	suite.Run(t, new(CommentSuite))
}

func (s *CommentSuite) SetupSuite() {
	conf, err := config.LoadConfig(".", ".test.env")
	if err != nil {
		s.Error(err)
	}

	s.db, s.dbmock, err = sqlmock.New()
	if err != nil {
		s.Error(err)
	}

	s.repo = repository.NewRepo(s.db)
	s.svc = service.GetServices(s.repo)
	s.ctrl = v1.NewController(s.svc, &conf)

	gin.SetMode(gin.ReleaseMode)
}

func (s *CommentSuite) TearDownSuite() {
	s.db.Close()
}

func (s *CommentSuite) TestCommentPost() {
	s.T().Run("OK", func(t *testing.T) {
		// Arrange
		userId := 1
		cmt := service.CommentBody{
			PostId:   2,
			ParentId: 0,
			Content:  "comment",
		}
		bodyBytes, _ := json.Marshal(cmt)
		body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("ID", userId)
		c.Request, _ = http.NewRequest(http.MethodPost, "/cmt", body)

		s.dbmock.ExpectExec("insert into Comment\\(userId, postId, parentId, content\\) values \\(\\$1, \\$2, \\$3, \\$4\\)").
			WithArgs(userId, cmt.PostId, cmt.ParentId, cmt.Content).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Act
		s.ctrl.PostComment(c)
		err := s.dbmock.ExpectationsWereMet()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	s.T().Run("PostId not found", func(t *testing.T) {
		// Arrange
		userId := 1
		cmt := service.CommentBody{
			PostId:   5,
			ParentId: 0,
			Content:  "comment",
		}
		bodyBytes, _ := json.Marshal(cmt)
		body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("ID", userId)
		c.Request, _ = http.NewRequest(http.MethodPost, "/cmt", body)

		s.dbmock.ExpectExec("insert into Comment\\(userId, postId, parentId, content\\) values \\(\\$1, \\$2, \\$3, \\$4\\)").
			WithArgs(userId, cmt.PostId, cmt.ParentId, cmt.Content).
			WillReturnError(errors.New(""))

		// Act
		s.ctrl.PostComment(c)
		err := s.dbmock.ExpectationsWereMet()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		//assert.Equal(t, `{"message":"2"}`, w.Body.String())
		assert.True(t, strings.HasPrefix(w.Body.String(), `{"message":`))
	})

	s.T().Run("Invalid body", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("ID", 1)
		// Act
		s.ctrl.PostComment(c)
		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Equal(t, `{"message":"Invalid json provided"}`, w.Body.String())
	})
}
