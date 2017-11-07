package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("gae-goa-sample", func() {
	Title("gae-goa-sample")
	Description("GAE goa サンプル")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("hello", func() {
	BasePath("/hello")

	Action("HelloWorld", func() {
		Description("HelloWorld")
		Routing(GET("/"))
		Response(OK, ExampleMedia)
		Response(NotFound)
	})
})

// ExampleMedia example
var ExampleMedia = MediaType("application/vnd.example+json", func() {
	Description("Example")
	Attributes(func() {
		Attribute("message", String, "message")
	})
	View("default", func() {
		Attribute("message")
	})
})
