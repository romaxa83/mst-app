package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/pkg/file"
	"io"
	"net/http"
	"os"
)

//var localizer *i18n.Localizer
//var bundle *i18n.Bundle

// @Summary Create author
// @Tags author
// @Description create authors of books
// @ID create-author
// @Accept  json
// @Produce  json
// @Param input body input.CreateAuthor true "author info"
// @Success 201 {object} resources.AuthorResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/authors [post]
func (h *Handler) createAuthor(c *gin.Context) {

	var input input.CreateAuthor
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.Author.Create(input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resources.NewAuthorResource(result))
}

// @Summary Get all authors paginator
// @Tags author
// @Description get all author with pagination data
// @ID get-all-author-pagination
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
// @Router /api/authors [get]
func (h *Handler) getAllAuthor(c *gin.Context) {
	var query input.GetAuthorQuery

	if err := c.Bind(&query); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	results, err := h.services.Author.GetAllPagination(query)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

// @Summary Get all authors list
// @Tags author
// @Description get all author list
// @ID get-all-author-list
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/authors/list [get]
func (h *Handler) getAllAuthorList(c *gin.Context) {

	results, err := h.services.Author.GetAllList()
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{results})
}

// @Summary Get one author
// @Tags author
// @Description get one author by id
// @ID get-one-author
// @Accept  json
// @Produce  json
// @Success 200 {object} resources.AuthorResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/authors/:id [get]
func (h *Handler) getOneAuthor(c *gin.Context) {

	result, err := h.services.Author.GetOne(getId(c))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resources.NewAuthorResource(result))
}

// @Summary Update author
// @Tags author
// @Description update author of books
// @ID update-author
// @Accept  json
// @Produce  json
// @Param input body input.UpdateAuthor true "author info"
// @Success 200 {object} resources.AuthorResource
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/authors/:id [put]
func (h *Handler) updateAuthor(c *gin.Context) {

	var input input.UpdateAuthor
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.Author.Update(getId(c), input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resources.NewAuthorResource(result))
}

// @Summary Delete author (soft)
// @Tags author
// @Description delete author of books (soft)
// @ID delete-author
// @Accept  json
// @Produce  json
// @Success 204 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/authors/:id [delete]
func (h *Handler) deleteAuthor(c *gin.Context) {

	if err := h.services.Author.Delete(getId(c)); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, response{"ok"})
}

// @Summary Upload image
// @Tags author
// @Description upload image to author
// @ID upload-img-author
// @Accept mpfd
// @Produce  json
// @Param file formData file true "file"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/authors/:id/upload [post]
func (h *Handler) uploadAuthor(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, models.MaxUploadSize)

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	if _, err := file.Read(buffer); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	contentType := http.DetectContentType(buffer)

	// Validate File Type
	if _, ex := models.ImageTypes[contentType]; !ex {
		errorResponse(c, http.StatusBadRequest, "file type is not supported")
		return
	}

	authorID := getId(c)

	// после загрузки файла на сторонний сервис, удаляем его
	tempFilename := fmt.Sprintf("%d-%s", authorID, fileHeader.Filename)

	f, err := os.OpenFile(tempFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "failed to create temp file")
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, bytes.NewReader(buffer)); err != nil {
		errorResponse(c, http.StatusInternalServerError, "failed to write chunk to temp file")
		return
	}

	_, err = h.services.Media.UploadAndSaveFile(c.Request.Context(), input.UploadMedia{
		OwnerID:     getId(c),
		OwnerType:   models.AuthorOwner,
		Type:        models.Image,
		ContentType: contentType,
		Name:        tempFilename,
		Size:        fileHeader.Size,
		Status:      models.ClientUpload,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"ok"})
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

// @Summary Import author
// @Tags author
// @Description import some authors
// @ID import-author
// @Accept mpfd
// @Produce  json
// @Param file formData file true "file"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/authors/import [post]
func (h *Handler) importAuthor(c *gin.Context) {

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, models.MaxUploadSize)

	fileReq, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer fileReq.Close()

	buffer := make([]byte, fileHeader.Size)
	if _, err := fileReq.Read(buffer); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	contentType := fileHeader.Header.Get("Content-Type")

	// Validate File Type
	if _, ex := models.ImportTypes[contentType]; !ex {
		errorResponse(c, http.StatusBadRequest,
			h.locale.GetTranslateWithReplace(
				getLocale(c),
				"file.invalid type",
				map[string]string{"Type": contentType}),
		)
		return
	}

	// после загрузки файла на сторонний сервис, удаляем его
	tempFilename := fmt.Sprintf("%s", fileHeader.Filename)
	//pathFilename := fmt.Sprintf("./%s/%s", "storage", tempFilename)

	f, err := os.OpenFile(tempFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError,
			h.locale.GetTranslate(getLocale(c), "failed to create temp file"))
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, bytes.NewReader(buffer)); err != nil {
		errorResponse(c, http.StatusInternalServerError,
			h.locale.GetTranslate(getLocale(c), "failed to write chunk to temp file"))
		return
	}

	defer file.Remove(tempFilename)
	m, err := h.services.Import.Create(input.CreateImport{
		Entity:      models.AuthorEntityImport,
		ContentType: models.ImportContentType(contentType),
		Status:      models.ImportUpload,
		FilePath:    tempFilename,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Import.Parse(m); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{h.locale.GetTranslate(getLocale(c), "message.ok")})
}

// @Summary Export author
// @Tags author
// @Description export authors
// @ID export-author
// @Produce  json
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /api/authors/export [get]
func (h *Handler) exportAuthor(c *gin.Context) {

	//path := "./"
	//fileName := c.Param("filename")
	//targetPath := filepath.Join(path, fileName)
	targetPath := "./authors.json"

	//logger.Warn(Exists(targetPath))

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	//c.Header("Content-Description", "File Transfer")
	//c.Header("Content-Transfer-Encoding", "binary")
	//c.Header("Content-Disposition", "attachment; filename=authors.json")
	//c.Header("Content-Type", "application/octet-stream")
	//c.File(targetPath)

	c.FileAttachment(targetPath, "authors.json")

	//c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, models.MaxUploadSize)
	//
	//file, fileHeader, err := c.Request.FormFile("file")
	//if err != nil {
	//	errorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	//defer file.Close()
	//
	//buffer := make([]byte, fileHeader.Size)
	//if _, err := file.Read(buffer); err != nil {
	//	errorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//contentType := http.DetectContentType(buffer)
	//
	//// Validate File Type
	//if _, ex := models.FileTypes[contentType]; !ex {
	//	errorResponse(c, http.StatusBadRequest, "file type is not supported")
	//	return
	//}
	//
	//// после загрузки файла на сторонний сервис, удаляем его
	//tempFilename := fmt.Sprintf("%s", fileHeader.Filename)
	//
	//f, err := os.OpenFile(tempFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	//if err != nil {
	//	errorResponse(c, http.StatusInternalServerError, "failed to create temp file")
	//	return
	//}
	//defer f.Close()
	//
	//if _, err := io.Copy(f, bytes.NewReader(buffer)); err != nil {
	//	errorResponse(c, http.StatusInternalServerError, "failed to write chunk to temp file")
	//	return
	//}

	//_, err = h.services.Media.UploadAndSaveFile(c.Request.Context(), input.UploadMedia{
	//	OwnerID:     getId(c),
	//	OwnerType:   models.AuthorOwner,
	//	Type:        models.Image,
	//	ContentType: contentType,
	//	Name:        tempFilename,
	//	Size:        fileHeader.Size,
	//	Status:      models.ClientUpload,
	//})
	//if err != nil {
	//	errorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}

	//c.JSON(http.StatusOK, response{"ok"})
}
