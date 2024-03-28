package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"ugmo/repository"

	"github.com/gin-gonic/gin"
)

func NewVideoHandler(videoService VideoService) videoHandler {
	return videoHandler{videoService: videoService}
}

func (h videoHandler) GetVideosId(c *gin.Context) {
	video := repository.Video{}
	err := c.ShouldBindJSON(&video)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	videos, err := h.videoService.GetVideosId(video)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, videos)
}

func (h videoHandler) GetImage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	path, err := h.videoService.GetImage(int(id))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Data(http.StatusOK, "image/jpeg", data)
}

func (h videoHandler) GetVideo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	path, err := h.videoService.GetVideo(int(id))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	ext := filepath.Ext(path)
	contentType := ""
	switch ext {
	case ".mp4":
		contentType = "video/mp4"
	default:
		c.String(http.StatusInternalServerError, "Unsupported file format")
		return
	}

	c.Data(http.StatusOK, contentType, data)
	// c.JSON(http.StatusOK, path)
}
