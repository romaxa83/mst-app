package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	languageHeader = "Content-language"
	localeCtx      = "locale"
)

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

// парсим заголовок и устанавливаем локаль в контекст
func (h *Handler) setLocale(c *gin.Context) {
	lang := c.GetHeader(languageHeader)

	c.Set(localeCtx, lang)
}

// получаем локаль из контекста
func getLocale(c *gin.Context) string {

	// todo обработать ошибку, а также установка дефолтного значения для языкаЮ если не был передан
	lang, _ := c.Get(localeCtx)

	return lang.(string)
}
