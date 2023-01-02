package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mager/penny-pincher/ctrl"
)

func ProvideRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/b/:id", ctrl.GetBudget)

	return r
}

var Options = ProvideRouter
