package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	services "blockchain/pkg/services"
)

type TransactionsController struct {
	TransactionsService services.TransactionsService
}

func NewTransactionsController(ts services.TransactionsService) TransactionsController {
	return TransactionsController{
		TransactionsService: ts,
	}
}

func (tc *TransactionsController) GetBalance(ctx *gin.Context) {
	address := ctx.Param("address")
	err, balance := tc.TransactionsService.GetBalance(&address)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": balance,
	})
}

func (tc *TransactionsController) Transfer(ctx *gin.Context) {
	transferParams := services.TransferParams{}
	if err := ctx.ShouldBind(&transferParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err, signedTx := tc.TransactionsService.Transfer(&transferParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Printf("transaction sent: %s", signedTx.Hash().Hex())
	ctx.JSON(http.StatusOK, gin.H{
		"data": signedTx.Hash().Hex(),
	})
}

func (tc *TransactionsController) QueryTransactions(ctx *gin.Context) {
	err, block := tc.TransactionsService.QueryTransactions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": block.Transactions(),
	})
}
