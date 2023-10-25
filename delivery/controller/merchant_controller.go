package controller

import (
	"go-transaction/model"
	"go-transaction/usecase"
	"go-transaction/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	router  *gin.Engine
	usecase usecase.MerchantUsecase
}

func (m *MerchantController) listHandler(ctx *gin.Context) {
	merchants, err := m.usecase.ShowListMerchant()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if len(merchants) == 0 {
		ctx.JSON(http.StatusNoContent, map[string]any{})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success show all merchants",
		"data":    merchants,
	})
}

func (m *MerchantController) createHandler(ctx *gin.Context) {
	var merchant model.Merchant
	merchant.Id = common.GenerateUUID()

	if err := ctx.ShouldBindJSON(&merchant); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if err := m.usecase.RegisterNewMerchant(merchant); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"status":  "succeed",
		"message": "success register new merchant",
	})
}

func (m *MerchantController) detailHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	merchant, err := m.usecase.GetDetailmerchant(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]any{
			"status":  "failed",
			"message": "data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success search merchant",
		"data":    merchant,
	})
}

func (m *MerchantController) updateHandler(ctx *gin.Context) {
	var merchant model.Merchant
	merchant.Id = ctx.Param("id")
	if err := ctx.ShouldBindJSON(&merchant); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if err := m.usecase.EditExistedMerchant(merchant); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success update existed merchant",
	})
}

func NewMerchantController(router *gin.Engine, usecase usecase.MerchantUsecase) *MerchantController {
	controller := &MerchantController{router, usecase}

	routerGroup := router.Group("/api/merchant")
	routerGroup.GET("/", controller.listHandler)
	routerGroup.POST("/", controller.createHandler)
	routerGroup.GET("/:id", controller.detailHandler)
	routerGroup.PUT("/:id", controller.updateHandler)

	return controller
}
