package router

import (
	"github.com/barlus-developer/go-simple-http/internal/infrastructure/config"
	"github.com/barlus-developer/go-simple-http/internal/interfaces/http/handler"
	"github.com/barlus-developer/go-simple-http/internal/interfaces/http/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func New(cfg config.Config, log *zap.Logger, healthHandler *handler.HealthHandler) *gin.Engine {
	if cfg.App.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger(log))

	engine.GET("/", healthHandler.Index)

	return engine
}
