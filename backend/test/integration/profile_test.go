//go:build integration

package integration

import (
	"math/rand"
	"strconv"
	"testing"

	"app/config"
	"app/internal/driver"
	"app/internal/model"
	"app/internal/repository"
	"app/internal/service"
	"app/pkg/auth"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// DB - Repo - Service
type ProfileSuite struct {
	suite.Suite
	repo        repository.ProfileRepo
	svc         service.ProfileService
	authManager *auth.Manager
}

func TestProfileSuite(t *testing.T) {
	suite.Run(t, new(ProfileSuite))
}

func (s *ProfileSuite) SetupSuite() {
	conf, err := config.LoadConfig(".", ".test.env")
	if err != nil {
		s.Error(err)
	}

	db := driver.Connect(conf.DbDriver, conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)
	err = db.SQL.Ping()
	if err != nil {
		s.Error(err)
	}

	s.repo = repository.NewProfileRepo(db.SQL)
	s.svc = service.NewProfileService(s.repo)

	s.authManager = auth.InitManager("id", conf.SigningKey)
}

func (s *ProfileSuite) TearDownSuite() {
}

func (s *ProfileSuite) TestProfileRegister() {
	s.T().Run("OK", func(t *testing.T) {
		// Arrange
		body := service.ProfileBody{
			Email:     strconv.Itoa(rand.Int()) + "@email.com",
			Username:  "username",
			Password:  "password",
			Gender:    "M",
			Birthdate: "2020-2-20",
		}

		// Act
		err1 := s.svc.Register(body)
		profile, err2 := s.svc.GetByEmail(body.Email)

		// Assert
		t.Log(profile)
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Equal(t, body.Username, profile.Name)
		assert.Equal(t, body.Gender, profile.Gender)
		assert.True(t, s.authManager.ComparePassword(body.Password, profile.Salt, profile.Hash))
	})

	s.T().Run("Email existed", func(t *testing.T) {
		// Arrange
		body := service.ProfileBody{
			Email:     "user1@gmail.com",
			Username:  "username",
			Password:  "password",
			Gender:    "M",
			Birthdate: "2020-2-20",
		}

		// Act
		err := s.svc.Register(body)

		// Assert
		assert.Error(t, err)
	})
}

func (s *ProfileSuite) TestProfileSelect() {
	s.T().Run("OK", func(t *testing.T) {
		// Arrange
		result := model.Profile{
			Id:     1,
			Name:   "User1",
			Gender: "M",
			Email:  "user1@gmail.com",
			Salt:   "SXmZdHRT",
			Hash:   "23h3nlI-gObbXQvg2DFDHClP4YA=",
		} // seed.up.sql

		// Act
		profile, err := s.svc.Get(result.Id)

		// Assert
		t.Log(profile)
		assert.NoError(t, err)
		assert.Equal(t, result.Id, profile.Id)
		assert.Equal(t, result.Name, profile.Name)
		assert.Equal(t, result.Gender, profile.Gender)
		assert.Equal(t, result.Email, profile.Email)
		assert.Equal(t, result.Hash, profile.Hash)
		assert.Equal(t, result.Salt, profile.Salt)
	})

	s.T().Run("Not found", func(t *testing.T) {
		// Arrange
		id := 0
		// Act
		profile, err := s.svc.Get(id)
		// Assert
		assert.Error(t, err)
		assert.Equal(t, 0, profile.Id)
		assert.Equal(t, "", profile.Email)
	})
}
