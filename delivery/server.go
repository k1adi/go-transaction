package delivery

import (
	"go-transaction/config"
	"go-transaction/delivery/controller"
	"go-transaction/manager"

	"fmt"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (a *appServer) initController() {
	controller.NewTestController(a.engine)
}

func (a *appServer) Run() {
	a.initController()

	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	engine := gin.Default()
	cfg := config.NewConfig()

	infraManager := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	host := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)
	return &appServer{
		engine:         engine,
		host:           host,
		usecaseManager: useCaseManager,
	}
}
