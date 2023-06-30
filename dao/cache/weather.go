package cache

import (
	"time"

	"github.com/apihutco/server/models"
)

const (
	weatherNowKeyPrefix  = "weather_now:"
	weatherNowExpiration = time.Minute * 10
)

type IWeather interface {
	Set(weather *models.Weather) error
	Get(location string) (*models.Weather, error)
}

var sharedWeatherCtrl *weatherCtrl

type weatherCtrl struct {
	cache *Cache
}

func (c *Cache) Weather() IWeather {
	if sharedWeatherCtrl == nil {
		sharedWeatherCtrl.cache = c
	}
	return sharedWeatherCtrl
}

func (w *weatherCtrl) Set(weather *models.Weather) error {
	return w.cache.redis.SetNX(getWeatherNowKey(weather.Location), weather, weatherNowExpiration).Err()
}

func (w *weatherCtrl) Get(location string) (*models.Weather, error) {
	weather := &models.Weather{}
	err := w.cache.redis.Get(getWeatherNowKey(location)).Scan(weather)
	return weather, err
}

func getWeatherNowKey(location string) string {
	return weatherNowKeyPrefix + location
}
