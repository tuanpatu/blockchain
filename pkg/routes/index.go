package routes

import (
	controllers "blockchain/pkg/controllers"
	models "blockchain/pkg/models"
	services "blockchain/pkg/services"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = models.ConnectDataBase()
	// jwtService service.JWTService = service.NewJWTService()

	usersService    services.UsersService
	usersController controllers.UsersController
	usersRoute      UsersRoute

	transactionsService    services.TransactionsService
	transactionsController controllers.TransactionsController
	transactionsRoute      TransactionsRoute

	walletsService    services.WalletsService
	walletsController controllers.WalletsController
	walletsRoute      WalletsRoute

	ctx context.Context
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func InitRouter() *gin.Engine {
	app := gin.Default()

	usersService = services.NewUsersService(db)
	usersController = controllers.NewUsersController(usersService)
	usersRoute = NewUsersRoute(usersController)

	transactionsService = services.NewTransactionsService(ctx)
	transactionsController = controllers.NewTransactionsController(transactionsService)
	transactionsRoute = NewTransactionsRoute(transactionsController)

	walletsService = services.NewWalletsService(db)
	walletsController = controllers.NewWalletsController(walletsService)
	walletsRoute = NewWalletRoute(walletsController)

	basepath := app.Group("/api")
	usersRoute.RegisterRoutes(basepath)
	transactionsRoute.RegisterRoutes(basepath)
	walletsRoute.RegisterRoutes(basepath)

	return app
}
