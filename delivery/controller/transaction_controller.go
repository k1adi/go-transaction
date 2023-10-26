package controller

import (
	"go-transaction/delivery/middleware"
	"go-transaction/model"
	"go-transaction/usecase"
	"go-transaction/utils/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	router  *gin.Engine
	usecase usecase.TransactionUsecase
}

func (t *TransactionController) createHandler(ctx *gin.Context) {
	var transaction model.Transaction
	transaction.Id = common.GenerateUUID()
	transaction.CustomerId = common.IDFromToken(ctx)
	transaction.TransactionAt = time.Now()

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if err := t.usecase.RegisterNewTransaction(transaction); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"status":  "succeed",
		"message": "success register new transaction",
	})
}

func NewTransactionController(router *gin.Engine, usecase usecase.TransactionUsecase) *TransactionController {
	controller := &TransactionController{router, usecase}

	routerGroup := router.Group("/api/transaction")
	routerGroup.POST("/", middleware.AuthMiddleware(), controller.createHandler)
	return controller
}
