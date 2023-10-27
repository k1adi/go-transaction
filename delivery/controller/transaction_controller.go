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

func (t *TransactionController) listHandler(ctx *gin.Context) {
	transactions, err := t.usecase.ShowListTransactions()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if len(transactions) == 0 {
		ctx.JSON(http.StatusNoContent, map[string]any{})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success show all transactions",
		"data":    transactions,
	})
}

func (t *TransactionController) historyHandler(ctx *gin.Context) {
	customerId := common.IDFromToken(ctx)
	transactions, err := t.usecase.ShowHistoryTransactions(customerId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if len(transactions) == 0 {
		ctx.JSON(http.StatusNoContent, map[string]any{})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "succeed",
		"message": "success show history transactions",
		"data":    transactions,
	})
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
	routerGroup.GET("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.listHandler)
	routerGroup.POST("/", middleware.AuthMiddleware(), middleware.UserMiddleware(), controller.createHandler)
	routerGroup.GET("/history", middleware.AuthMiddleware(), middleware.UserMiddleware(), controller.historyHandler)
	return controller
}
