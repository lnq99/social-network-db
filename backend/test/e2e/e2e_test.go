//go:build e2e

package e2e

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/config"
	v1 "app/internal/controller/v1"
	"app/internal/driver"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type E2ESuite struct {
	suite.Suite
	db   *driver.DB
	repo *repository.Repo
	svc  *service.Services
	ctrl *v1.Controller
}

func TestE2ESuite(t *testing.T) {
	suite.Run(t, new(E2ESuite))
}

func (s *E2ESuite) SetupSuite() {
	conf, err := config.LoadConfig(".", ".test.env")
	if err != nil {
		s.Error(err)
	}

	s.db = driver.Connect(conf.DbDriver, conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)
	err = s.db.SQL.Ping()
	if err != nil {
		s.Error(err)
	}

	s.repo = repository.NewRepo(s.db.SQL)
	s.svc = service.GetServices(s.repo)
	s.ctrl = v1.NewController(s.svc, &conf)

	gin.SetMode(gin.ReleaseMode)

	//router := controller.NewRouter()
	//router = s.ctrl.SetupRouter(router)
	//router.Run(":" + conf.Port)
}

func (s *E2ESuite) TearDownSuite() {
	//s.db.SQL.Close()
}

func (s *E2ESuite) TestCommentPost() {
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

		// Act
		s.ctrl.PostComment(c)
		//resp, err := http.Post(cmtUrl, "application/json", bytes.NewBuffer(bodyBytes))

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())
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

		// Act
		s.ctrl.PostComment(c)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.NotEqual(t, "", w.Body.String())
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

func BenchmarkE2E(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//
	}
}
