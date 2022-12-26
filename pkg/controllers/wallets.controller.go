package controllers

import (
	services "blockchain/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletsController struct {
	WalletsService services.WalletsService
}

func NewWalletsController(ws services.WalletsService) WalletsController {
	return WalletsController{
		WalletsService: ws,
	}
}

func (wc *WalletsController) GetWallet(ctx *gin.Context) {
	id := ctx.Param("id")
	err, Wallet := wc.WalletsService.GetWallet(&id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Wallet fetched succesfully",
		"data":    Wallet,
	})
}
