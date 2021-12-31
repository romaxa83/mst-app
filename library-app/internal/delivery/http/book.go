package http

import (
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"net/http"
)

// @Summary Create book
// @Tags book
// @Description create book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param input body input.CreateBook true "book info"
// @Success 201 {object} resources.BookResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/books [post]
func (h *Handler) createBook(c *gin.Context) {

	var input input.CreateBook
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//logger.Warnf("%+v", input)

	result, err := h.services.Book.Create(input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resources.NewBookResource(result))
}

// @Summary Get one book
// @Tags book
// @Description get one book by id
// @ID get-one-book
// @Accept  json
// @Produce  json
// @Success 200 {object} resources.BookResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/books/:id [get]
func (h *Handler) getOneBook(c *gin.Context) {

	result, err := h.services.Book.GetOne(getId(c))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resources.NewBookResource(result))
}

// @Summary Get all books paginator
// @Tags book
// @Description get all books with pagination data
// @ID get-all-book-pagination
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param sort query string false "sort"
// @Param search query string false "search"
// @Param id query int false "id"
// @Success 200 {object} db.Pagination
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/books [get]
func (h *Handler) getAllBook(c *gin.Context) {
	var query input.GetBookQuery

	if err := c.Bind(&query); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	results, err := h.services.Book.GetAllPagination(query)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

// @Summary Update book
// @Tags book
// @Description update book
// @ID update-book
// @Accept  json
// @Produce  json
// @Param input body input.UpdateBook true "book info"
// @Success 200 {object} resources.BookResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/books/:id [put]
func (h *Handler) updateBook(c *gin.Context) {

	var input input.UpdateBook
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.Book.Update(getId(c), input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resources.NewBookResource(result))
}

// @Summary Delete book (soft)
// @Tags book
// @Description delete book
// @ID delete-book
// @Accept  json
// @Produce  json
// @Success 204 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/books/:id [delete]
func (h *Handler) deleteBook(c *gin.Context) {

	if err := h.services.Book.Delete(getId(c)); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, response{"ok"})
}
