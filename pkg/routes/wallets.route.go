package routes

import (
	"github.com/gin-gonic/gin"

	controllers "blockchain/pkg/controllers"
)

type WalletsRoute struct {
	WalletsController controllers.WalletsController
}

func NewWalletRoute(WalletsController controllers.WalletsController) WalletsRoute {
	return WalletsRoute{
		WalletsController: WalletsController,
	}
}

func (wr *WalletsRoute) RegisterRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/wallets")
	{
		routes.GET("/:id", wr.WalletsController.GetWallet)
	}
}
