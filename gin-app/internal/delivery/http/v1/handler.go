package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/gin-app/internal/services"
	"github.com/romaxa83/mst-app/gin-app/pkg/auth"
)

type Handler struct {
	Services     *services.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *services.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		Services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
		h.initListRoutes(v1)
		h.initItemRoutes(v1)
		h.initUploadRoutes(v1)
	}
}
