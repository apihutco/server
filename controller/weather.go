package controller

import (
	"apihut-server/logic/weather"
	"apihut-server/response"

	"github.com/gin-gonic/gin"
)

func WeatherNowHandler(c *gin.Context) {
	location := c.Query("location")

	info, err := weather.GetNowWeather(location)
	if err != nil {
		response.BadRequest(c).JSON()
	}

	response.Success(c).Data(info).JSON()
}

func WeatherDay3Handler(c *gin.Context) {
	location := c.Query("location")

	info, err := weather.GetDay3Weather(location)
	if err != nil {
		response.BadRequest(c).JSON()
	}

	response.Success(c).Data(info).JSON()
}
