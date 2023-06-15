package controller

import (
	"apihut-server/logic/geo_bank"
	"apihut-server/response"

	"github.com/gin-gonic/gin"
)

func GeoBankHandler(c *gin.Context) {
	location := c.Query("location")
	adm := c.Query("adm")

	info, err := geo_bank.GetGeoInfo(location, adm)
	if err != nil {
		response.BadRequest(c).Code(response.CodeError).JSON()
		return
	}

	response.Success(c).Data(info).JSON()
}
