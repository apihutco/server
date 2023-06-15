package weather

import (
	"errors"

	"apihut-server/models"
)

var weatherCtrlList []IWeatherCtrl

type IWeatherCtrl interface {
	Now(location string) (*models.Weather, error)
	Day3(location string) ([]*models.Weather, error)
}

func Init() {
	weatherCtrlList = make([]IWeatherCtrl, 0)
	weatherCtrlList = append(weatherCtrlList, NewQWeather())
}

func GetNowWeather(location string) (*models.Weather, error) {
	for _, ctrl := range weatherCtrlList {
		weatherInfo, err := ctrl.Now(location)
		if err != nil {
			return nil, err
		}
		return weatherInfo, nil
	}

	return nil, errors.New("Now weather get error")
}

func GetDay3Weather(location string) ([]*models.Weather, error) {
	for _, ctrl := range weatherCtrlList {
		weatherInfo, err := ctrl.Day3(location)
		if err != nil {
			return nil, err
		}
		return weatherInfo, nil
	}

	return nil, errors.New("Now weather get error")
}
