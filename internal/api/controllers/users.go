package controllers

import (
	"net/http"

	"github.com/vanbien2402/first-web-demo/internal/api/domain"
	"github.com/vanbien2402/first-web-demo/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService domain.UserService
}

//Register Register User function
func (ctl *userController) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.CreateParams
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	resp, err := ctl.userService.Create(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &resp)
}

//Login user log in
func (ctl *userController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	userName, password, ok := c.Request.BasicAuth()
	if len(userName) == 0 || len(password) == 0 || !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user name or password invalid"})
	}
	user, err := ctl.userService.Validate(ctx, userName, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	tokenString, err := jwt.GenerateJWT(user.UserName, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

//GetUser get user
func (ctl *userController) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.GetParams
	if err := c.BindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	resp, err := ctl.userService.Get(ctx, req.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &resp)
}

//UpdateUser Update User function
func (ctl *userController) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.UpdateParams
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if req.ID != c.Param("id") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id mismatch"})
		c.Abort()
		return
	}
	resp, err := ctl.userService.Update(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &resp)
}

//DeleteUser Delete User function
func (ctl *userController) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.DeleteParams
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	err := ctl.userService.Delete(ctx, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, "Delete success")
}

//NewUsersController init UserController
func NewUsersController(userService domain.UserService) domain.UserController {
	return &userController{
		userService: userService,
	}
}
