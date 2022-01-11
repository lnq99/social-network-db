package main

import (
	"log"
	"os"

	"app/config"
	"app/internal/controller"
	v1 "app/internal/controller/v1"
	"app/internal/driver"
	"app/internal/repository"
	"app/internal/service"
	// _ "net/http/pprof"
	// /debug/pprof
)

// @title Social Network
// @version 1.0
// @description Tiny social network.

// @contact.name Quang Le
// @contact.email lenhuquang99@gmail.com

// @license.name Apache 2.0

// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	confFile := ".env"
	if len(os.Args) == 2 {
		confFile = os.Args[1]
		log.Println(confFile)
	}

	var err error
	conf, err := config.LoadConfig(".", confFile)
	if err != nil {
		panic(err)
	}

	// log.Println(conf.DbDriver)

	db := driver.Connect(conf.DbDriver, conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)
	err = db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepo(db.SQL)
	services := service.GetServices(repo)

	// test.RepoSelect(repo)

	ctrl := v1.NewController(services, &conf)

	router := controller.NewRouter()
	router = ctrl.SetupRouter(router)
	router = controller.SwaggerRouter(router)

	addr := conf.Host + ":" + conf.Port
	router.Run(addr)
}
