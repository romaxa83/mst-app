package http

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/library-app/internal/config"
	"github.com/romaxa83/mst-app/library-app/internal/services"
	"github.com/romaxa83/mst-app/library-app/internal/utils"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "github.com/romaxa83/mst-app/library-app/docs"
)

// dependency injection

type Handler struct {
	services *services.Services
	locale   *utils.Local
}

// construct

func NewHandler(services *services.Services, locale *utils.Local) *Handler {
	return &Handler{
		services: services,
		locale:   locale,
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
		author := api.Group("/authors", h.setLocale)
		{
			author.POST("/", h.createAuthor)
			author.GET("/", h.getAllAuthor)
			author.GET("/list", h.getAllAuthorList)
			author.GET("/:id", h.getOneAuthor)
			author.PUT("/:id", h.updateAuthor)
			author.DELETE("/:id", h.deleteAuthor)
			author.POST("/:id/upload", h.uploadAuthor)
			author.POST("/import", h.importAuthor)
			author.GET("/export", h.exportAuthor)
		}
		book := api.Group("/books")
		{
			book.POST("/", h.createBook)
			book.GET("/", h.getAllBook)
			book.GET("/:id", h.getOneBook)
			book.PUT("/:id", h.updateBook)
			book.DELETE("/:id", h.deleteBook)
		}

		archive := api.Group("/archive")
		{
			archive.GET("/categories", h.archiveCategory)
			archive.PUT("/categories/restore/:id", h.restoreCategory)
			archive.DELETE("/categories/:id", h.deleteCategoryForce)
		}
	}

	return router
}
