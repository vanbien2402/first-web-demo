package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vanbien2402/first-web-demo/internal/api/configs"
	"github.com/vanbien2402/first-web-demo/internal/api/controllers"
	"gorm.io/gorm"
)

type route struct {
	config          *configs.Config
	usersController controllers.IUsersController
}

func initRouter(router *gin.Engine,
	config *configs.Config,
	db *gorm.DB) {
	controllers.NewUsersController(db)
	r := &route{
		config:          config,
		usersController: controllers.NewUsersController(db),
	}
	r.initUsersRouter(router)
}

func (r *route) initUsersRouter(router *gin.Engine) {
	router.GET("/users", r.usersController.GetUsers)
	router.POST("/users", r.usersController.CreateUser)
	router.PUT("/users/:id", r.usersController.UpdateUser)
	router.DELETE("/users/:id", r.usersController.DeleteUser)
}
