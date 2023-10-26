package delivery

import (
	"go-transaction/config"
	"go-transaction/delivery/controller"
	"go-transaction/delivery/middleware"
	"go-transaction/manager"

	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type appServer struct {
	usecaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (a *appServer) initController() {
	storeCookie := cookie.NewStore([]byte("secret"))

	a.engine.Use(sessions.Sessions("session", storeCookie))
	a.engine.Use(middleware.LoggerMiddleware(a.log))
	controller.NewBankController(a.engine, a.usecaseManager.BankUsecase())
	controller.NewMerchantController(a.engine, a.usecaseManager.MerchantUsecase())
	controller.NewCustomerController(a.engine, a.usecaseManager.CustomerUsecase(), a.usecaseManager.AuthUsecase())
	controller.NewAdminController(a.engine, a.usecaseManager.AdminUsecase(), a.usecaseManager.AuthUsecase())
	controller.NewTransactionController(a.engine, a.usecaseManager.TransactionUsecase())
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
		log:            logrus.New(),
	}
}
