package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/vanbien2402/first-web-demo/internal/api/configs"
	"github.com/vanbien2402/first-web-demo/internal/pkg/rds"
)

type Server struct {
	router *gin.Engine
}

//Start server
func (s *Server) Start() error {
	return s.router.Run(":8080")
}

//NewServer initialize new server
func NewServer(config *configs.Config) (*Server, error) {
	db, err := rds.Connect(&config.DB)
	if err != nil {
		log.Fatal("connect DB failed", err)
		return nil, err
	}
	router := gin.New()
	_ = router.SetTrustedProxies(nil)
	initRouter(router, config, db)
	return &Server{
		router: router,
	}, nil
}
