package domain

import (
	"context"

	"github.com/vanbien2402/first-web-demo/internal/api/models"
	"github.com/vanbien2402/first-web-demo/internal/pkg/rds"

	"github.com/gin-gonic/gin"
)

//UserController User Controller interface
type UserController interface {
	//Register Register User function
	Register(c *gin.Context)
	//Login log in
	Login(c *gin.Context)
	//GetUser GetByID User function
	GetUser(c *gin.Context)
	//UpdateUser Update User function
	UpdateUser(c *gin.Context)
	//DeleteUser Delete User function
	DeleteUser(c *gin.Context)
}

//UserService User Service interface
type UserService interface {
	//Validate validate log in user
	Validate(ctx context.Context, userName, password string) (res *UserInfo, err error)
	//Get with params
	Get(ctx context.Context, userName string) (res *UserInfo, err error)
	//Create with params
	Create(ctx context.Context, req CreateParams) (user *models.User, err error)
	//Update with params
	Update(ctx context.Context, req UpdateParams) (user *models.User, err error)
	//Delete with params
	Delete(ctx context.Context, id string) (err error)
}

//UserRepository User Repository interface
type UserRepository interface {
	//FindByUserName with params
	FindByUserName(ctx context.Context, userName string) (*models.User, error)
	//Create with params
	Create(ctx context.Context, user *models.User) error
	//Update with params
	Update(ctx context.Context, user *models.User) error
	//Delete with params
	Delete(ctx context.Context, id string) error
	//CheckExist check register user name is exist
	CheckExist(ctx context.Context, userName string) bool
	//FindByID with params
	FindByID(ctx context.Context, id string) (*models.User, error)
}

//GetParams query users params
type GetParams struct {
	UserName string `json:"userName"`
}

//CreateParams create user params
type CreateParams struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

//UpdateParams update user params
type UpdateParams struct {
	ID       string `json:"id" binding:"required"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Version  int64  `json:"version" binding:"required"`
}

//DeleteParams delete user params
type DeleteParams struct {
	ID string `json:"id"`
}

type UserInfo struct {
	rds.Model
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
