package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getId(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return 0
	}

	return id
}
