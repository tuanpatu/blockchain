package routes

import (
	"github.com/gin-gonic/gin"

	controllers "blockchain/pkg/controllers"
)

type TransactionsRoute struct {
	TransactionsController controllers.TransactionsController
}

func NewTransactionsRoute(TransactionsController controllers.TransactionsController) TransactionsRoute {
	return TransactionsRoute{
		TransactionsController: TransactionsController,
	}
}

func (tr *TransactionsRoute) RegisterRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/transactions")
	{
		routes.POST("", tr.TransactionsController.Transfer)
		routes.GET("", tr.TransactionsController.QueryTransactions)
		routes.GET("/:address", tr.TransactionsController.GetBalance)
	}
}
