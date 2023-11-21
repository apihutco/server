package controller

import (
	"fmt"
	"net/http"

	"github.com/apihutco/server/models"
	"go.uber.org/zap"

	"github.com/apihutco/server/config"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	path := c.Request.URL.Path
	msg := fmt.Sprintf("API Not Found. \nYou path is: %s , check it again please. \nOur docs: %s", path, config.Conf.Site.DocsUrl)
	c.String(http.StatusOK, msg)
}

func VersionHandler(c *gin.Context) {
	zap.L().Debug("apihut info",
		zap.String("version", config.VERSION),
		zap.String("build_time", config.BUILD_TIME),
	)

	c.JSON(http.StatusOK, &models.AppInfo{
		Version:   config.VERSION,
		BuildTime: config.BUILD_TIME,
	})
}
