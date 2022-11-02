package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "APIHut")
}

func NotFound(c *gin.Context) {
	c.String(http.StatusOK, "Not found")
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "true"})
}
