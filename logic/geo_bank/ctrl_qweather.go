package geo_bank

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"apihut-server/config"
	"apihut-server/models"
	"apihut-server/utils/consts"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type qWeather struct {
	GeoBase
	key string
}

type qWeatherLocationRsp struct {
	Code     string `json:"code"`
	Location []struct {
		Name      string `json:"name"`
		Id        string `json:"id"`
		Lat       string `json:"lat"`
		Lon       string `json:"lon"`
		Adm2      string `json:"adm2"`
		Adm1      string `json:"adm1"`
		Country   string `json:"country"`
		Tz        string `json:"tz"`
		UtcOffset string `json:"utcOffset"`
		IsDst     string `json:"isDst"`
		Type      string `json:"type"`
		Rank      string `json:"rank"`
		FxLink    string `json:"fxLink"`
	} `json:"location"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}

func (q *qWeather) GetInfo(location string, adm string) (*models.GeoBank, error) {
	const baseURL = "https://geoapi.qweather.com/v2/city/lookup"
	v := url.Values{}
	v.Set("key", q.key)
	v.Set("location", location)
	v.Set("adm", adm)

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()

	zap.L().Debug("QWeatherLocation", zap.String("URL", u.String()))

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

	locationInfo := &qWeatherLocationRsp{}
	if err = json.Unmarshal(body, locationInfo); err != nil {
		return nil, err
	}

	if locationInfo.Code != consts.QWeatherLocationSuccessCode {
		return nil, errors.New(locationInfo.Code)
	}

	if len(locationInfo.Location) < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	l := locationInfo.Location[0]
	geoID, _ := strconv.Atoi(l.Id)

	return &models.GeoBank{
		Country:   l.Country,
		Province:  l.Adm1,
		City:      l.Adm2,
		District:  l.Name,
		GeoID:     int64(geoID),
		TimeZone:  l.Tz,
		UTCOffset: l.UtcOffset,
	}, nil
}

func NewQWeather() IGeoCtrl {
	return &qWeather{
		key: config.Conf.QWeather.Key,
	}
}
