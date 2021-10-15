package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sample-go-app/cmd/api/handlers/getuser"
	"github.com/sample-go-app/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users", getuser.Do(app))
	return mux
}
