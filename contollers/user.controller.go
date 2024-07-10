package contollers

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
	ctx.JSON(http.StatusOK, gin.H{"Message": "message"})
}


func (userCtrl *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "message"})
}


func (userCtrl *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "message"})
}


func (userCtrl *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "message"})
}

// Route group for user controller
func (userCtrl *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userroute := rg.Group("/api/user")
	userroute.POST("/create", userCtrl.CreateUser)
	userroute.GET("/:name", userCtrl.GetUser)
	userroute.GET("/detail", userCtrl.GetAll)
	userroute.PUT("/update", userCtrl.UpdateUser)
	userroute.DELETE("/delete", userCtrl.DeleteUser)
}



