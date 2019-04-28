package transformer

import "github.com/fasthttp/router"

func NewRoutes() *router.Router {
	r := router.New()

	r.OPTIONS("/transform", TransformHandle)
	r.POST("/transform", TransformHandle)

	return r
}
