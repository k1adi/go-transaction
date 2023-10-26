package controller

import (
	"go-transaction/delivery/middleware"
	"go-transaction/model"
	"go-transaction/usecase"
	"go-transaction/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router      *gin.Engine
	usecase     usecase.CustomerUsecase
	authUsecase usecase.AuthUsecase
}

func (c *CustomerController) listHandler(ctx *gin.Context) {
	customers, err := c.usecase.ShowListCustomers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if len(customers) == 0 {
		ctx.JSON(http.StatusNoContent, map[string]any{})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success show all merchants",
		"data":    customers,
	})
}

func (c *CustomerController) loginHandler(ctx *gin.Context) {
	var account model.Auth
	account.Role = "user"
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	token, err := c.authUsecase.Login(account, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"status":  "succeed",
		"message": token,
	})
}

func (c *CustomerController) registerHandler(ctx *gin.Context) {
	var customer model.Customer
	customer.Id = common.GenerateUUID()
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if err := c.usecase.RegisterNewCustomer(customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"status":  "succeed",
		"message": "success register new customer",
	})
}

func (c *CustomerController) logoutHandler(ctx *gin.Context) {
	c.authUsecase.Logout(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "succeed",
		"message": "success logout",
	})
}

func NewCustomerController(router *gin.Engine, usecase usecase.CustomerUsecase, authUsecase usecase.AuthUsecase) *CustomerController {
	controller := &CustomerController{router, usecase, authUsecase}

	routerGroup := router.Group("/api/customer")
	routerGroup.GET("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.listHandler)
	routerGroup.POST("/login", controller.loginHandler)
	routerGroup.POST("/register", controller.registerHandler)
	routerGroup.POST("/logout", controller.logoutHandler)

	return controller
}
