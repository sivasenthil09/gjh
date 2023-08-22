package controllers

import (
	"net/http"
	"strings"
	"rest-api/interfaces"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService interfaces.ITransaction
}

func InitTransactionController(transactionService interfaces.ITransaction) TransactionController {
	return TransactionController{transactionService} //DI(dependency injection) pattern
}

func  (pc *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction *models.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := pc.TransactionService.CreateTransaction(transaction)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}

func (pc *ProfileController)  GetTransactionById(context *gin.Context)     {}

