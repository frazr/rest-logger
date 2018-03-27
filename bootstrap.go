package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type RequestMethod func(string, ...gin.HandlerFunc) gin.IRoutes

func (r *RestLogger) Init() {
	r.setRequestMethods()
	r.setRoutes()
	r.setResponses()
	r.initRoutes()
}

func (r *RestLogger) setRequestMethods() {
	r.RequestMethods = make(map[string]RequestMethod)
	r.RequestMethods["POST"] = r.Engine.POST
	r.RequestMethods["GET"] = r.Engine.GET
	r.RequestMethods["DELETE"] = r.Engine.DELETE
}

func (r *RestLogger) setRoutes() {
	r.Routes.SetPostRoutes()
	r.Routes.SetGetRoutes()
}

func (r *RestLogger) setResponses() {
	r.SuccessResponse = gin.H{"message": "Successful request"}
	r.ErrorResponse = gin.H{"message": "Erroneus request"}
}

func (r *RestLogger) initRoutes() {
	for _, route := range r.Routes {
		f, ok := r.RequestMethods[route.Type]
		if ok {
			f(route.URI, func(c *gin.Context) {
				c.JSON(200, r.SuccessResponse)
			})
		} else {
			log.Print("Request Method not implemented/available")
		}
	}
}
