//go:generate goagen bootstrap -d github.com/y-ogura/gae-goa-sample/design

package server

import (
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/y-ogura/gae-goa-sample/app"
	"github.com/y-ogura/gae-goa-sample/controller"
)

func init() {
	// Create service
	service := goa.New("gae-goa-sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "hello" controller
	c := controller.NewHelloController(service)
	app.MountHelloController(service, c)

	// setup Handler
	http.HandleFunc("/", service.Mux.ServeHTTP)
}
