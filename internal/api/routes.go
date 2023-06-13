package api

import (
	"github.com/vanbien2402/first-web-demo/internal/api/configs"
	"github.com/vanbien2402/first-web-demo/internal/api/controllers"
	"github.com/vanbien2402/first-web-demo/internal/api/domain"
	"github.com/vanbien2402/first-web-demo/internal/api/middleware"
	"github.com/vanbien2402/first-web-demo/internal/api/repositories"
	"github.com/vanbien2402/first-web-demo/internal/api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type route struct {
	config          *configs.Config
	usersController domain.UserController
}

func initRouter(router *gin.Engine,
	config *configs.Config,
	db *gorm.DB) {
	//init repository
	userRepo := repositories.NewUsersRepository(db)

	//init service
	userService := services.NewUserService(userRepo)

	r := &route{
		config:          config,
		usersController: controllers.NewUsersController(userService),
	}
	api := router.Group("/api")
	r.initCommonRouter(api)
	secured := api.Group("/secured")
	r.initSecuredRouter(secured)
}

func (r *route) initCommonRouter(group *gin.RouterGroup) {
	group.POST("/user/register", r.usersController.Register)
	group.POST("/user/login", r.usersController.Login)
	group.GET("/users", r.usersController.GetUser)
	//group.PUT("/users/:id", r.usersController.UpdateUser)
}

func (r *route) initSecuredRouter(g *gin.RouterGroup) {
	group := g.Use(middleware.Auth())
	{
		group.PUT("/users/:id", r.usersController.UpdateUser)
		group.DELETE("/users/:id", r.usersController.DeleteUser)
	}
}
