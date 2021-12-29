package http

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/library-app/internal/config"
	"github.com/romaxa83/mst-app/library-app/internal/services"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "github.com/romaxa83/mst-app/library-app/docs"
)

// dependency injection

type Handler struct {
	services *services.Services
}

// construct

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

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

	api := router.Group("/api")
	{
		category := api.Group("/categories")
		{
			category.POST("/", h.createCategory)
			category.GET("/", h.getAllCategory)
			category.GET("/list", h.getAllCategoryList)
			category.GET("/:id", h.getOneCategory)
			category.PUT("/:id", h.updateCategory)
			category.DELETE("/:id", h.deleteCategory)
		}

		archive := api.Group("/archive")
		{
			archive.GET("/categories", h.archiveCategory)
			archive.PUT("/categories/restore/:id", h.restoreCategory)
		}
	}

	return router
}
