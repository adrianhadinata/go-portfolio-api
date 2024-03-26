package delivery

import (
	"fmt"
	"portfolio-api/config"
	"portfolio-api/delivery/controller"
	"portfolio-api/manager"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ucManager manager.UsecaseManager
	engine *gin.Engine
	config *config.Config
	host   string
}

func (s *Server) loadController() {
	apiGroup :=s.engine.Group("/api") 
	v1 := apiGroup.Group("/v1")

	controller.NewVisitsController(s.ucManager.VisitsUsecase(), v1).Route()
		
}

func Start (s *Server) {
	s.loadController()
	err := s.engine.Run(s.host)

	if err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	config, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	infraManager, err := manager.NewInfraManager(config)

	if err != nil {
		panic(err)
	}

	repoManager := manager.NewRepoManager(infraManager)
	ucManager := manager.NewUseCaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", config.ApiConfig.ApiPort)

	return &Server {
		ucManager: ucManager,
		engine: engine,
		config: config,
		host: host,
	}
}