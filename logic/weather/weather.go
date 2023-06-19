package weather

import (
	"errors"

	"apihut-server/dao/redis"
	"apihut-server/models"

	"go.uber.org/zap"
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

func GetNowWeather(location string) (weatherInfo *models.Weather, err error) {
	// 从缓存中加载
	weatherInfo, err = redis.GetWeatherCache(location)
	if err != nil {
		zap.L().Error("从缓存中加载实时天气数据失败", zap.Error(err))
	}
	if err == nil && weatherInfo != nil {
		zap.L().Debug("从缓存中加载实时天气数据成功", zap.Any("info", weatherInfo))
		return weatherInfo, nil
	}

	for _, ctrl := range weatherCtrlList {
		weatherInfo, err = ctrl.Now(location)
		if err != nil {
			return nil, err
		}
	}

	if weatherInfo != nil {
		zap.L().Debug("设置实时天气缓存", zap.Any("details", weatherInfo))
		if err = redis.SetWeatherCache(weatherInfo); err != nil {
			zap.L().Error("设置实时天气缓存失败", zap.Error(err))
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
