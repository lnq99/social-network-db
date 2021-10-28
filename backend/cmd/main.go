package main

import (
	"app/config"
	"app/internal/controller"
	v1 "app/internal/controller/v1"
	"app/internal/driver"
	"app/internal/repository"
	"log"
)

func main() {
	var err error
	conf, err := config.LoadConfig(".", ".env")
	if err != nil {
		panic(err)
	}

	log.Println(conf.DbDriver)

	db := driver.Connect(conf.DbDriver, conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)
	err = db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepo(db.SQL)

	// test.RepoSelect(repo)

	ctrl := v1.NewController(&repo, &conf)

	router := controller.NewRouter()
	router = ctrl.SetupRouter(router)
	router = controller.SwaggerRouter(router)

	router.Run()
}
