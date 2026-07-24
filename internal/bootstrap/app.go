package bootstrap

import (
	"github.com/barlus-developer/go-simple-http/internal/application/health"
	"github.com/barlus-developer/go-simple-http/internal/infrastructure/config"
	"github.com/barlus-developer/go-simple-http/internal/infrastructure/logger"
	httpHandler "github.com/barlus-developer/go-simple-http/internal/interfaces/http/handler"
	httpRouter "github.com/barlus-developer/go-simple-http/internal/interfaces/http/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	Config config.Config
	Logger *zap.Logger
	Router *gin.Engine
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	log, err := logger.New(cfg.App.Environment)
	if err != nil {
		return nil, err
	}

	healthService := health.NewService()
	healthHandler := httpHandler.NewHealthHandler(healthService)
	router := httpRouter.New(cfg, log, healthHandler)

	return &App{
		Config: cfg,
		Logger: log,
		Router: router,
	}, nil
}
