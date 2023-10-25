package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	router *gin.Engine
}

func (a *TestController) testHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"message": "success create test handler",
	})
}

func NewTestController(router *gin.Engine) {
	controller := &TestController{
		router: router,
	}

	routerGroup := router.Group("/api/test")
	routerGroup.GET("/", controller.testHandler)
}
