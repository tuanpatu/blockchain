package routes

import (
	"github.com/gin-gonic/gin"

	controllers "blockchain/pkg/controllers"
)

type UsersRoute struct {
	UsersController controllers.UsersController
}

func NewUsersRoute(Userscontroller controllers.UsersController) UsersRoute {
	return UsersRoute{
		UsersController: Userscontroller,
	}
}

func (ur *UsersRoute) RegisterRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/users")
	{
		routes.POST("", ur.UsersController.CreateUser)
		routes.GET("", ur.UsersController.GetUsers)
		routes.GET("/:id", ur.UsersController.GetUser)
		routes.PUT("/:id", ur.UsersController.UpdateUser)
		routes.DELETE("/:id", ur.UsersController.DeleteUser)
	}
}
