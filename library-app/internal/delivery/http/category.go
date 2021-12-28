package http

import (
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	"net/http"
	"strconv"
)

// @Summary Create category
// @Tags category
// @Description create categories of books
// @ID create-category
// @Accept  json
// @Produce  json
// @Param input body input.CreateCategory true "category info"
// @Success 201 {object} resources.CategoryResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/categories [post]
func (h *Handler) createCategory(c *gin.Context) {

	var input input.CreateCategory
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.Category.Create(input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resources.NewCategoryResource(result))
}

// @Summary Get all categories paginator
// @Tags category
// @Description get all categories with pagination data
// @ID get-all-category-pagination
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Success 200 {object} db.Pagination
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/categories [get]
func (h *Handler) getAllCategory(c *gin.Context) {
	var query db.Pagination
	if err := c.Bind(&query); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//logger.Infof("%+v", query)

	results, err := h.services.Category.GetAllPagination(query)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

type getAllListsResponse struct {
	Rows interface{} `json:"rows"`
}

// @Summary Get all categories list
// @Tags category
// @Description get all categories list
// @ID get-all-category-list
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/categories/list [get]
func (h *Handler) getAllCategoryList(c *gin.Context) {

	results, err := h.services.Category.GetAllList()
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{results})
}

// @Summary Get one category
// @Tags category
// @Description get one category by id
// @ID get-one-category
// @Accept  json
// @Produce  json
// @Success 200 {object} resources.CategoryResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/categories/:id [get]
func (h *Handler) getOneCategory(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	result, err := h.services.Category.GetOne(id)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resources.NewCategoryResource(result))
}

// @Summary Update category
// @Tags category
// @Description update category of books
// @ID update-category
// @Accept  json
// @Produce  json
// @Param input body input.UpdateCategory true "category info"
// @Success 200 {object} resources.CategoryResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/categories/:id [put]
func (h *Handler) updateCategory(c *gin.Context) {

	var input input.UpdateCategory
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	result, err := h.services.Category.Update(id, input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resources.NewCategoryResource(result))
}

// @Summary Delete category (soft)
// @Tags category
// @Description delete category of books (soft)
// @ID delete-category
// @Accept  json
// @Produce  json
// @Success 204 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/categories/:id [delete]
func (h *Handler) deleteCategory(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Category.Delete(id); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, response{"ok"})
}
