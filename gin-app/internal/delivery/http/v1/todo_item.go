package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"net/http"
	"strconv"
)

func (h *Handler) initItemRoutes(api *gin.RouterGroup) {
	lists := api.Group("/lists", h.userIdentity)
	{
		items := lists.Group(":id/items")
		{
			items.GET("/", h.getAllItems)
			items.POST("/", h.createItem)
		}
	}

	items := api.Group("items", h.userIdentity)
	{
		items.GET("/:id", h.getItemById)
		items.PUT("/:id", h.updateItem)
		items.DELETE("/:id", h.deleteItem)
	}
}

// @Summary Create Item
// @Security ApiKeyAuth
// @Tags todo.txt-items
// @Description create items for todo.txt-list
// @ID create-item
// @Accept  json
// @Produce  json
// @Param input body domains.TodoItem true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/lists/:id/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var input domains.TodoItem
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.TodoItem.Create(c.Request.Context(), userId, listId, input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllItemsResponse struct {
	Data []domains.TodoItem `json:"data"`
}

// @Summary Get All Items
// @Security ApiKeyAuth
// @Tags todo.txt-items
// @Description get all items by list
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllItemsResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid lis id param")
		return
	}

	items, err := h.Services.TodoItem.GetAll(c.Request.Context(), userId, listId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

// @Summary Get One Item
// @Security ApiKeyAuth
// @Tags todo.txt-items
// @Description get one item by id
// @ID get-one-items
// @Accept  json
// @Produce  json
// @Success 200 {object} domains.TodoItem
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/items/:id [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.Services.TodoItem.GetById(c.Request.Context(), userId, id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update Item
// @Security ApiKeyAuth
// @Tags todo.txt-items
// @Description update item
// @ID update-item
// @Accept  json
// @Produce  json
// @Param input body domains.UpdateItemInput true "update item info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/items/:id [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domains.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Services.TodoItem.Update(userId, id, input); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"ok"})
}

// @Summary Delete Item (soft)
// @Security ApiKeyAuth
// @Tags todo.txt-items
// @Description delete item (soft)
// @ID delete-list
// @Accept  json
// @Produce  json
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/v1/items/:id [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.Services.TodoItem.Delete(userId, itemId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"ok"})
}
