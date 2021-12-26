package v1

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"

	"io"
	"net/http"
	"os"
)

const (
	maxUploadSize = 5 << 20 // 5 megabytes
	maxVideoSize  = 2 << 30 // 2 gigabytes
)

var (
	imageTypes = map[string]interface{}{
		"image/jpeg": nil,
		"image/png":  nil,
	}

	videoTypes = map[string]interface{}{
		"video/mp4":                 nil,
		"application/octet-stream":  nil,
		"text/plain; charset=utf-8": nil, // for strange files with such content-type
	}

	fileTypes = map[string]interface{}{
		"application/pdf":           nil,
		"application/zip":           nil, // excel
		"text/plain; charset=utf-8": nil, // for strange files with such content-type
	}
)

type uploadResponse struct {
	URL string `json:"url"`
}

func (h *Handler) initUploadRoutes(api *gin.RouterGroup) {
	upload := api.Group("/upload")
	{
		upload.POST("/image", h.uploadImage)
		upload.POST("/file", h.uploadFile)
	}
}

// @Summary Upload Image
// @Tags upload
// @Description upload image
// @ModuleID uploadImage
// @Accept mpfd
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} uploadResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/upload/image [post]
func (h *Handler) uploadImage(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

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
	if _, ex := imageTypes[contentType]; !ex {
		errorResponse(c, http.StatusBadRequest, "file type is not supported")
		return
	}

	// после загрузки файла на сторонний сервис, удаляем его
	tempFilename := fmt.Sprintf("%s", fileHeader.Filename)

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

	url, err := h.Services.Files.UploadAndSaveFile(c.Request.Context(), domains.File{
		Type:        domains.Image,
		ContentType: contentType,
		Name:        tempFilename,
		Size:        fileHeader.Size,
	})

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &uploadResponse{url})
}

// @Summary Upload File
// @Tags upload
// @Description upload file
// @ModuleID uploadFile
// @Accept mpfd
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} uploadResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/upload/image [post]
func (h *Handler) uploadFile(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

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

	//// Validate File Type
	if _, ex := fileTypes[contentType]; !ex {
		errorResponse(c, http.StatusBadRequest, "file type is not supported")

		return
	}

	// после загрузки файла на сторонний сервис, удаляем его
	tempFilename := fmt.Sprintf("%s", fileHeader.Filename)

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

	url, err := h.Services.Files.UploadAndSaveFile(c.Request.Context(), domains.File{
		Type:        domains.Other,
		ContentType: contentType,
		Name:        tempFilename,
		Size:        fileHeader.Size,
	})

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &uploadResponse{url})
}
