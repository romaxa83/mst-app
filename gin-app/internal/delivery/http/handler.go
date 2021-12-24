package http

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/gin-app/internal/config"
	v1 "github.com/romaxa83/mst-app/gin-app/internal/delivery/http/v1"
	"github.com/romaxa83/mst-app/gin-app/internal/services"
	"github.com/romaxa83/mst-app/gin-app/pkg/auth"
	"net/http"

	//"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/romaxa83/mst-app/gin-app/docs"
)

// dependency injection

type Handler struct {
	services     *services.Services
	tokenManager auth.TokenManager
}

// construct

func NewHandler(services *services.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

//func prometheusHandler() gin.HandlerFunc {
//	h := promhttp.Handler()
//	return func(c *gin.Context) {
//		h.ServeHTTP(c.Writer, c.Request)
//	}
//}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	if cfg.Environment != config.Prod {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
