package controller

import (
	"go-transaction/model"
	"go-transaction/usecase"
	"go-transaction/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	router  *gin.Engine
	usecase usecase.BankUsecase
}

func (b *BankController) listHandler(ctx *gin.Context) {
	banks, err := b.usecase.ShowListBank()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if len(banks) == 0 {
		ctx.JSON(http.StatusNoContent, map[string]any{})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success show all banks",
		"data":    banks,
	})
}

func (b *BankController) createHandler(ctx *gin.Context) {
	var bank model.Bank
	bank.Id = common.GenerateUUID()

	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if err := b.usecase.RegisterNewBank(bank); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"status":  "succeed",
		"message": "success register new bank",
	})
}

func (b *BankController) detailHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	bank, err := b.usecase.GetDetailBank(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]any{
			"status":  "failed",
			"message": "data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success search bank",
		"data":    bank,
	})
}

func (b *BankController) updateHandler(ctx *gin.Context) {
	var bank model.Bank
	bank.Id = ctx.Param("id")
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if err := b.usecase.EditExistedBank(bank); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success update existed bank",
	})
}

func NewBankController(router *gin.Engine, usecase usecase.BankUsecase) *BankController {
	controller := &BankController{router, usecase}

	routerGroup := router.Group("/api/bank")
	routerGroup.GET("/", controller.listHandler)
	routerGroup.POST("/", controller.createHandler)
	routerGroup.GET("/:id", controller.detailHandler)
	routerGroup.PUT("/:id", controller.updateHandler)

	return controller
}
