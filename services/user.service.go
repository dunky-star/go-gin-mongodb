package services

import "dunky-star/gin-mongodb-apis/models"

//  Services contract for CRUD operations on User struct
type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
    GetAll() ([] *models.User, error)
    UpdateUser(*models.User) error
	DeleteUser(*string) error
}