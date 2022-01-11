//go:build e2e_bdd

package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"testing"

	"app/config"
	"app/internal/controller"
	v1 "app/internal/controller/v1"
	"app/internal/driver"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBddSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bdd Suite")
}

var (
	db                 *driver.DB
	registerUrl        string
	server             http.Server
	fixtureProfileBody = []service.ProfileBody{
		service.ProfileBody{
			Email:     "user1@gmail.com",
			Username:  "username",
			Password:  "password",
			Gender:    "M",
			Birthdate: "2020-2-20",
		},
	}
)

var _ = BeforeSuite(func() {
	conf, err := config.LoadConfig(".", ".test.env")
	Expect(err).To(BeNil())

	db = driver.Connect(conf.DbDriver, conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)
	err = db.SQL.Ping()
	Expect(err).To(BeNil())

	registerUrl = "http://localhost:" + conf.Port + "/api/v1/auth/register"

	repo := repository.NewRepo(db.SQL)
	svc := service.GetServices(repo)
	ctrl := v1.NewController(svc, &conf)
	Expect(ctrl).NotTo(BeNil())

	gin.SetMode(gin.ReleaseMode)

	router := controller.NewRouter()
	router = ctrl.SetupRouter(router)
	//go router.Run(":" + conf.Port)
	server = http.Server{
		Addr:    ":" + conf.Port,
		Handler: router,
	}

	addr := server.Addr
	ln, err := net.Listen("tcp", addr)
	Expect(err).To(BeNil())

	log.Printf("Server listening on: %s\n", addr)
	go func() {
		if err = server.Serve(ln); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("%s\n", err)
		}
	}()
})

var _ = AfterSuite(func() {
	server.Shutdown(context.Background())
	//db.SQL.Close()
})

var _ = Describe("Posting comment", Label("Comment"), func() {

	When("request has body", func() {
		Context("email not existed", func() {
			It("register new account", func() {
				body := service.ProfileBody{
					Email:     strconv.Itoa(rand.Int()) + "@email.com",
					Username:  "username",
					Password:  "password",
					Gender:    "M",
					Birthdate: "2020-2-20",
				}
				bodyBytes, _ := json.Marshal(body)

				resp, err := http.Post(registerUrl, "application/json", bytes.NewBuffer(bodyBytes))

				Expect(resp).NotTo(BeNil())
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(http.StatusCreated))
			})

		})

		Context("email existed", func() {
			for _, body := range fixtureProfileBody {
				It("return error account existed", func() {
					bodyBytes, _ := json.Marshal(body)

					resp, err := http.Post(registerUrl, "application/json", bytes.NewBuffer(bodyBytes))

					Expect(resp).NotTo(BeNil())
					Expect(err).To(BeNil())
					Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
				})
			}
		})
	})

	When("request has no body", func() {
		It("should get error response", func() {
			resp, err := http.Post(registerUrl, "application/json", nil)

			Expect(resp).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusUnprocessableEntity))
		})
	})
})
