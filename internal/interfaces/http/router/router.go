package router

import (
	"net/http"

	"github.com/barlus-developer/go-simple-http/internal/infrastructure/config"
	"github.com/barlus-developer/go-simple-http/internal/interfaces/http/handler"
	"github.com/barlus-developer/go-simple-http/internal/interfaces/http/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func New(cfg config.Config, log *zap.Logger, healthHandler *handler.HealthHandler) *gin.Engine {
	if !cfg.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.HandleMethodNotAllowed = true
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger(log))

	engine.GET("/", healthHandler.Index)

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "The requested resource was not found.",
		})
	})

	engine.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"status":  "error",
			"message": "The requested method is not allowed for this resource.",
		})
	})

	return engine
}
