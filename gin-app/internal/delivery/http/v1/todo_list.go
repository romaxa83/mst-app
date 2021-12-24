package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"net/http"
	"strconv"
)

func (h *Handler) initListRoutes(api *gin.RouterGroup) {
	lists := api.Group("/lists", h.userIdentity)
	{
		lists.GET("/", h.getAll)
		lists.POST("/", h.create)
		lists.GET("/:id", h.getById)
		lists.PUT("/:id", h.update)
		lists.DELETE("/:id", h.delete)
	}
}

// @Summary Create List
// @Security ApiKeyAuth
// @Tags todo-lists
// @Description create list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body domains.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/lists [post]
func (h *Handler) create(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input domains.TodoList
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.TodoList.Create(c.Request.Context(), userId, input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []domains.TodoList `json:"data"`
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags todo-lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/lists [get]
func (h *Handler) getAll(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.Services.TodoList.GetAll(c.Request.Context(), userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get One List
// @Security ApiKeyAuth
// @Tags todo-lists
// @Description get one list by id
// @ID get-one-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} domains.TodoList
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/lists/:id [get]
func (h *Handler) getById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.Services.TodoList.GetById(c.Request.Context(), userId, id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update List
// @Security ApiKeyAuth
// @Tags todo-lists
// @Description update list
// @ID update-list
// @Accept  json
// @Produce  json
// @Param input body domains.UpdateTodoListInput true "update list info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/lists/:id [put]
func (h *Handler) update(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domains.UpdateTodoListInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Services.TodoList.Update(c.Request.Context(), userId, id, input); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"ok"})
}

// @Summary Delete List (soft)
// @Security ApiKeyAuth
// @Tags todo-lists
// @Description delete list (soft)
// @ID delete-list
// @Accept  json
// @Produce  json
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/lists/:id [delete]
func (h *Handler) delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.Services.TodoList.Delete(c.Request.Context(), userId, id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"ok"})
}
