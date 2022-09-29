package logic

import (
	"apihut-server/config"
	"apihut-server/logic/consts"
	"apihut-server/models"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

type tencent struct {
	key string
}

type tencentRsp struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Result  struct { // IP 定位结果
		Ip       string `json:"ip"` // IP
		Location struct {
			Lat float64 `json:"lat"` // 纬度
			Lng float64 `json:"lng"` // 经度
		} `json:"location"`
		AdInfo struct { // 定位行政区划信息
			Nation   string `json:"nation"`   // 国家
			Province string `json:"province"` // 省份
			City     string `json:"city"`     // 市
			District string `json:"district"` // 区
			Adcode   int    `json:"adcode"`   // 行政区划代码
		} `json:"ad_info"`
	} `json:"result"`
}

func TencentInit(key string) IIPCtrl {
	return &tencent{key: key}
}

func (t *tencent) GetIP(ip net.IP) (*models.IPBank, error) {
	v := url.Values{}
	v.Set("key", config.Share.Open.Tencent.Key)
	v.Set("ip", ip.String())

	u, err := url.Parse("https://apis.map.qq.com/ws/location/v1/ip")
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	tUrl := u.String()

	resp, err := http.Get(tUrl)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	lbs := &tencentRsp{}
	if err = json.Unmarshal(body, lbs); err != nil {
		return nil, err
	}

	if lbs.Status != 0 {
		return nil, errors.New("")
	}

	return &models.IPBank{
		IP:       ip,
		Country:  lbs.Result.AdInfo.Nation,
		Province: lbs.Result.AdInfo.Province,
		City:     lbs.Result.AdInfo.City,
		District: lbs.Result.AdInfo.District,
		ISP:      "",
		Location: ToLocation(lbs.Result.Location.Lat, lbs.Result.Location.Lng),
		Source:   t.Platform().Name(),
	}, nil
}

func (t *tencent) Platform() consts.PlatformID {
	return consts.Tencent
}
