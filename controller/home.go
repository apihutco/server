package controller

import (
	"fmt"
	"net/http"

	"apihut-server/config"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	path := c.Request.URL.Path
	msg := fmt.Sprintf("API Not Found. \nYou path is: %s , check it again please. \nOur docs: %s", path, config.Conf.Site.DocsUrl)
	c.String(http.StatusOK, msg)
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "true"})
}
