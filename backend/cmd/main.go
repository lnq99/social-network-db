package main

import (
	"app/config"
	"app/internal/driver"
	"app/internal/repository"
	"app/test"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile(conf.StaticRoot, true)))

	r.Run()
}
