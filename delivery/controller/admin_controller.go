package controller

import (
	"go-transaction/delivery/middleware"
	"go-transaction/model"
	"go-transaction/usecase"
	"go-transaction/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	router      *gin.Engine
	usecase     usecase.AdminUsecase
	authUsecase usecase.AuthUsecase
}

func (a *AdminController) loginHandler(ctx *gin.Context) {
	var account model.Auth
	account.Role = "admin"
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	token, err := a.authUsecase.Login(account, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "admin",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
		MaxAge:   300,
	})

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": token,
	})
}

func (a *AdminController) createHandler(ctx *gin.Context) {
	var admin model.Auth
	admin.Id = common.GenerateUUID()
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if err := a.usecase.RegisterNewAdmin(admin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"status":  "succeed",
		"message": "success register new admin",
	})
}

func (a *AdminController) listHandler(ctx *gin.Context) {
	admins, err := a.usecase.ShowListAdmins()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if len(admins) == 0 {
		ctx.JSON(http.StatusNoContent, map[string]any{})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success show all admins",
		"data":    admins,
	})
}

func (a *AdminController) logoutHandler(ctx *gin.Context) {
	a.authUsecase.Logout(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "succeed",
		"message": "success logout",
	})
}

func NewAdminController(router *gin.Engine, usecase usecase.AdminUsecase, authUsecase usecase.AuthUsecase) *AdminController {
	controller := &AdminController{router, usecase, authUsecase}

	routerGroup := router.Group("/api/admin")
	routerGroup.POST("/login", controller.loginHandler)
	routerGroup.POST("/logout", controller.logoutHandler)
	routerGroup.POST("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.createHandler)
	routerGroup.GET("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.listHandler)

	return controller
}
