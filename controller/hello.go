package controller

import (
	"github.com/goadesign/goa"
	"github.com/y-ogura/gae-goa-sample/app"
)

// HelloController implements the hello resource.
type HelloController struct {
	*goa.Controller
}

// NewHelloController creates a hello controller.
func NewHelloController(service *goa.Service) *HelloController {
	return &HelloController{Controller: service.NewController("HelloController")}
}

// HelloWorld runs the HelloWorld action.
func (c *HelloController) HelloWorld(ctx *app.HelloWorldHelloContext) error {
	// HelloController_HelloWorld: start_implement

	// Put your logic here

	// HelloController_HelloWorld: end_implement
	msg := "Hello, World!"
	res := &app.Example{&msg}
	return ctx.OK(res)
}
