package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sample-go-app/cmd/api/router"
	"github.com/sample-go-app/pkg/application"
	"github.com/sample-go-app/pkg/exithandler"
	"github.com/sample-go-app/pkg/logger"
	"github.com/sample-go-app/pkg/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := server.
		Get().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(router.Get(app)).
		WithErrLogger(logger.Error)

	go func() {
		logger.Info.Printf("starting server at %s", app.Cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}
