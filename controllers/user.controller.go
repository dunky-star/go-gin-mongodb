package controllers

import (
	"dunky-star/gin-mongodb-apis/models"
	"dunky-star/gin-mongodb-apis/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService // Injecting UserService into controller
}


// Constructor 
func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (userCtrl *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
    err := userCtrl.UserService.CreateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "User created successfully!"})
}

func (userCtrl *UserController) GetUser(ctx *gin.Context) {
	var username = ctx.Param("name")
	user, err := userCtrl.UserService.GetUser(&username)
 	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}


func (userCtrl *UserController) GetAll(ctx *gin.Context) {
	users, err := userCtrl.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}


func (userCtrl *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := userCtrl.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}


func (userCtrl *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := userCtrl.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "User deleted successfully!"})
}

// Route group for user controller
func (userCtrl *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userroute := rg.Group("/v1")
	userroute.POST("/users", userCtrl.CreateUser)
	userroute.GET("/users/:name", userCtrl.GetUser)
	userroute.GET("/users", userCtrl.GetAll)
	userroute.PUT("/users/:name", userCtrl.UpdateUser)
	userroute.DELETE("/users/:name", userCtrl.DeleteUser)
}



