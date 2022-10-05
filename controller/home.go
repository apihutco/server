package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Home")
}

func NotFound(c *gin.Context) {
	c.String(http.StatusOK, "Not found")
}
