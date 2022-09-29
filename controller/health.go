package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "true"})
}
