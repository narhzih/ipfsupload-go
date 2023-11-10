package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api *Api) loadUploadRoutes() {
	api.BaseRouteGroup.POST("/upload", api.uploadHandler)
}

func (api *Api) uploadHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All good",
	})
}
