package controller

import (
	"apihut-server/logger"
	"encoding/base64"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
	"time"

	"apihut-server/logic/avatar"
	"apihut-server/logic/consts"
	"apihut-server/models"
	"apihut-server/response"

	"github.com/gin-gonic/gin"
)

/*
兼容 Gravatar，参考文档：
https://en.gravatar.com/site/implement/images/
*/

func AvatarHandler(c *gin.Context) {
	hash := c.Param("hash")
	if len(hash) == 0 {
		hash = c.ClientIP()
	}

	req := models.NewAvatar(hash)

	// 只要出现random参数，不管值是什么，都开启随机模式
	if _, has := c.GetQuery("random"); has {
		hash = strconv.Itoa(int(time.Now().UnixNano()))
		req.Hash = hash
		req.Random = true
	}

	err := c.ShouldBindQuery(&req)
	if err != nil {
		logger.L().Debug("参数绑定失败", zap.Error(err), zap.Any("query", c.Request.URL.RawQuery))
		response.ErrorWithCode(c, response.ErrorBind)
		return
	}

	filePath, err := avatar.NewAvatar(&req)
	if err != nil {
		logger.L().Error("生成头像失败", zap.Error(err))
		response.ErrorWithCode(c, response.ErrorAvatarGenerate)
		return
	}

	// 按格式返回
	switch req.GetOutput() {
	case consts.JSON:
		b := FileToBase64(filePath)
		response.SuccessWithData(c, gin.H{"avatar": b})
		return
	case consts.Base64:
		b := FileToBase64(filePath)
		c.String(http.StatusOK, b)
		return
	default:
		c.File(filePath)
		return
	}
}

/*
base64 格式图片格式

data:image/gif;base64,base64编码的gif图片数据
data:image/png;base64,base64编码的png图片数据
data:image/jpeg;base64,base64编码的jpeg图片数据
data:image/x-icon;base64,base64编码的icon图片数据
*/

func FileToBase64(name string) string {
	f, _ := os.Open(name)
	defer f.Close()

	fileInfo, _ := f.Stat()
	fileSize := fileInfo.Size()

	buffer := make([]byte, fileSize)
	_, _ = f.Read(buffer)

	// ext := strings.ReplaceAll(path.Ext(name), ".", "")
	// imageExtMap := map[string]struct{}{"png": {}, "jpg": {}, "jpeg": {}}
	//
	// b := base64.StdEncoding.EncodeToString(buffer)
	// if e, ok := imageExtMap[ext]; ok {
	// 	return fmt.Sprintf("data:image/%s;base64,%s", e, b)
	// } else {
	// 	return fmt.Sprintf("data:image/jpeg;base64,%s", b)
	// }

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buffer)
}
