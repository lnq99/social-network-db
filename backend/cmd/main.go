package main

import (
	"app/config"
	"app/internal/controller"
	"app/internal/driver"
	"app/internal/repository"
	"app/test"
)

func main() {
	var err error
	conf, err := config.LoadConfig(".", ".env")
	if err != nil {
		panic(err)
	}

	db := driver.Connect(conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)
	err = db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepo(db.SQL)

	test.RepoSelect(repo)

	ctrl := controller.Controller{Repo: &repo, Conf: &conf}

	r := controller.SetupRouter(ctrl)

	r.Run()
}
