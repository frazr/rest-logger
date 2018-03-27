package main

import (
	"github.com/gin-gonic/gin"
)

type RestLogger struct {
	Engine          *gin.Engine
	Routes          Routes
	SuccessResponse gin.H
	ErrorResponse   gin.H
	RequestMethods  map[string]RequestMethod
}

func main() {
	r := RestLogger{
		Engine: gin.Default(),
	}

	r.Init()
	r.Engine.Run()
}
