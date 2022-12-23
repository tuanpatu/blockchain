package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blockchain/pkg/models"
	services "blockchain/pkg/services"
)

type UsersController struct {
	UsersService services.UsersService
}

func NewUsersController(ts services.UsersService) UsersController {
	return UsersController{
		UsersService: ts,
	}
}

func (tc *UsersController) GetUsers(ctx *gin.Context) {
	err, Users := tc.UsersService.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Users fetched succesfully",
		"data":    Users,
	})
}

func (tc *UsersController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err, User := tc.UsersService.GetUser(&id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User fetched succesfully",
		"data":    User,
	})
}

func (tc *UsersController) CreateUser(ctx *gin.Context) {
	User := models.Users{}

	if err := ctx.BindJSON(&User); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err, resUser := tc.UsersService.CreateUser(&User)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User add succesfully",
		"data":    resUser,
	})
}

func (tc *UsersController) UpdateUser(ctx *gin.Context) {
	User := models.Users{}
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "id is null",
		})
		return
	}

	if err := ctx.BindJSON(&User); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err, resUser := tc.UsersService.UpdateUser(&User, &id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "could not update User",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User update succesfully",
		"data":    resUser,
	})
}

func (tc *UsersController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "id is null",
		})
		return
	}

	err := tc.UsersService.DeleteUser(&id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "could not delete User",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User delete succesfully",
	})
}
