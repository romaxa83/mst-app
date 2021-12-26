package http

import (
	"github.com/gin-gonic/gin"
)

type inputCreateCategory struct {
	title string
}

// @Summary Create category
// @Tags category
// @Description create categories of books
// @ID create-category
// @Accept  json
// @Produce  json
// @Param input body inputCreateCategory true "category info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/categories [post]
func (h *Handler) createCategory(c *gin.Context) {
	//userId, err := getUserId(c)
	//if err != nil {
	//	return
	//}
	//
	//var input domains.TodoList
	//if err := c.BindJSON(&input); err != nil {
	//	errorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//id, err := h.Services.TodoList.Create(c.Request.Context(), userId, input)
	//if err != nil {
	//	errorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"id": id,
	//})
}
