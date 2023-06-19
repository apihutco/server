package redis

import (
	"time"

	"apihut-server/models"
)

const (
	weatherNowKeyPrefix  = "weather_now:"
	weatherNowExpiration = time.Minute * 10
)

func SetWeatherCache(weather *models.Weather) error {
	return client.SetNX(getWeatherNowKey(weather.Location), weather, weatherNowExpiration).Err()
}

func GetWeatherCache(location string) (*models.Weather, error) {
	weather := &models.Weather{}
	err := client.Get(getWeatherNowKey(location)).Scan(weather)
	return weather, err
}

func getWeatherNowKey(location string) string {
	return weatherNowKeyPrefix + location
}
