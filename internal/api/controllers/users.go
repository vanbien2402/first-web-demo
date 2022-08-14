package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vanbien2402/first-web-demo/internal/models"
	"gorm.io/gorm"
	"log"
)

type usersController struct {
	db *gorm.DB
}

//IUsersController UsersController interface
type IUsersController interface {
	//GetUsers Get Users function
	GetUsers(c *gin.Context)
	//CreateUser Create User function
	CreateUser(c *gin.Context)
	//UpdateUser Update User function
	UpdateUser(c *gin.Context)
	//DeleteUser Delete User function
	DeleteUser(c *gin.Context)
}

//GetUsers Get Users function
func (ctl *usersController) GetUsers(c *gin.Context) {
	var users []models.User
	ctl.db.Find(&users)
	c.JSON(200, &users)
}

//CreateUser Create User function
func (ctl *usersController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("bind json failed when create user", err)
	}
	ctl.db.Create(&user)
	c.JSON(200, user)
}

//UpdateUser Update User function
func (ctl *usersController) UpdateUser(c *gin.Context) {
	var user models.User
	ctl.db.Where("id = ?", c.Param("id")).First(&user)
	if err := c.BindJSON(&user); err != nil {
		log.Println("bind json failed when update user", err)
	}
	ctl.db.Save(&user)
	c.JSON(200, &user)
}

//DeleteUser Delete User function
func (ctl *usersController) DeleteUser(c *gin.Context) {
	var user models.User
	ctl.db.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, "Delete success")
}

//NewUsersController init UsersController
func NewUsersController(db *gorm.DB) IUsersController {
	return &usersController{
		db: db,
	}
}
