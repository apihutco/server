package weather

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/apihutco/server/config"
	"github.com/apihutco/server/models"
	"github.com/apihutco/server/utils/consts"
)

// QWeather 和风天气
type QWeather struct {
	key string
}

type qWeatherNowResp struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Now        struct {
		ObsTime   string `json:"obsTime"`
		Temp      string `json:"temp"`
		FeelsLike string `json:"feelsLike"`
		Icon      string `json:"icon"`
		Text      string `json:"text"`
		Wind360   string `json:"wind360"`
		WindDir   string `json:"windDir"`
		WindScale string `json:"windScale"`
		WindSpeed string `json:"windSpeed"`
		Humidity  string `json:"humidity"`
		Precip    string `json:"precip"`
		Pressure  string `json:"pressure"`
		Vis       string `json:"vis"`
		Cloud     string `json:"cloud"`
		Dew       string `json:"dew"`
	} `json:"now"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}

type qWeatherDay3Resp struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Daily      []struct {
		FxDate         string `json:"fxDate"`
		Sunrise        string `json:"sunrise"`
		Sunset         string `json:"sunset"`
		Moonrise       string `json:"moonrise"`
		Moonset        string `json:"moonset"`
		MoonPhase      string `json:"moonPhase"`
		MoonPhaseIcon  string `json:"moonPhaseIcon"`
		TempMax        string `json:"tempMax"`
		TempMin        string `json:"tempMin"`
		IconDay        string `json:"iconDay"`
		TextDay        string `json:"textDay"`
		IconNight      string `json:"iconNight"`
		TextNight      string `json:"textNight"`
		Wind360Day     string `json:"wind360Day"`
		WindDirDay     string `json:"windDirDay"`
		WindScaleDay   string `json:"windScaleDay"`
		WindSpeedDay   string `json:"windSpeedDay"`
		Wind360Night   string `json:"wind360Night"`
		WindDirNight   string `json:"windDirNight"`
		WindScaleNight string `json:"windScaleNight"`
		WindSpeedNight string `json:"windSpeedNight"`
		Humidity       string `json:"humidity"`
		Precip         string `json:"precip"`
		Pressure       string `json:"pressure"`
		Vis            string `json:"vis"`
		Cloud          string `json:"cloud"`
		UvIndex        string `json:"uvIndex"`
	} `json:"daily"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}

func (q *QWeather) Now(location string) (*models.Weather, error) {
	const baseUrl = "https://devapi.qweather.com/v7/weather/now"

	v := url.Values{}
	v.Set("location", location)
	v.Set("key", q.key)

	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	u.RawQuery = v.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	weatherNow := new(qWeatherNowResp)
	if err = json.Unmarshal(body, weatherNow); err != nil {
		return nil, err
	}

	if weatherNow.Code != consts.QWeatherLocationSuccessCode {
		return nil, errors.New(weatherNow.Code)
	}

	return &models.Weather{
		Location:  location,
		ObsTime:   weatherNow.Now.ObsTime,
		Temp:      weatherNow.Now.Temp,
		FeelsLike: weatherNow.Now.FeelsLike,
		Text:      weatherNow.Now.Text,
		Wind360:   weatherNow.Now.Wind360,
		WindDir:   weatherNow.Now.WindDir,
		Humidity:  weatherNow.Now.Humidity,
		Precip:    weatherNow.Now.Precip,
		FxLink:    weatherNow.FxLink,
		Sources:   weatherNow.Refer.Sources,
	}, nil
}

func (q *QWeather) Day3(location string) ([]*models.Weather, error) {
	const baseUrl = "https://devapi.qweather.com/v7/weather/3d"

	v := url.Values{}
	v.Set("location", location)
	v.Set("key", q.key)

	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	u.RawQuery = v.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	weatherDay3 := new(qWeatherDay3Resp)
	if err = json.Unmarshal(body, weatherDay3); err != nil {
		return nil, err
	}

	if weatherDay3.Code != consts.QWeatherLocationSuccessCode {
		return nil, errors.New(weatherDay3.Code)
	}

	finalDay3Weather := make([]*models.Weather, 0)
	for _, dayItem := range weatherDay3.Daily {

		finalDay3Weather = append(finalDay3Weather, &models.Weather{
			Location:  location,
			FxDate:    dayItem.FxDate,
			MinTemp:   dayItem.TempMin,
			MaxTemp:   dayItem.TempMax,
			TextDay:   dayItem.TextDay,
			TextNight: dayItem.TextNight,
			Wind360:   dayItem.Wind360Day,
			WindDir:   dayItem.WindDirDay,
			Humidity:  dayItem.Humidity,
			Precip:    dayItem.Precip,
			FxLink:    weatherDay3.FxLink,
		})
	}

	return finalDay3Weather, nil
}

func NewQWeather() IWeatherCtrl {
	return &QWeather{
		key: config.Conf.QWeather.Key,
	}
}
