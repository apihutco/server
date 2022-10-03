package ip_bank

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"apihut-server/config"
	"apihut-server/logic/consts"
	"apihut-server/models"

	"github.com/pkg/errors"
)

type gaode struct {
	key string
}

type gaodeRsp struct {
	Status   string `json:"status,omitempty"`
	Info     string `json:"info,omitempty"`
	Infocode string `json:"infocode,omitempty"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	ISP      string `json:"isp"`
	Location string `json:"location"`
	IP       string `json:"ip"`
}

func GaodeInit() IIPCtrl {
	return &gaode{key: config.Share.Open.Gaode.Key}
}

func (g *gaode) GetIP(ip net.IP) (*models.IPBank, error) {
	v := url.Values{}
	v.Set("key", config.Share.Open.Gaode.Key)
	v.Set("ip", ip.String())
	v.Set("type", strconv.Itoa(IPVersion(ip.String())))
	v.Set("parameters", "")

	u, err := url.Parse("https://restapi.amap.com/v5/ip")
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	gUrl := u.String()

	resp, err := http.Get(gUrl)
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

	lbs := &gaodeRsp{}
	if err = json.Unmarshal(body, lbs); err != nil {
		return nil, err
	}

	if lbs.Status != "1" {
		return nil, errors.New(lbs.Info)
	}

	return &models.IPBank{
		IP:       ip.String(),
		Country:  lbs.Country,
		Province: lbs.Province,
		City:     lbs.City,
		District: lbs.District,
		ISP:      lbs.ISP,
		Location: lbs.Location,
		Source:   g.Platform().Name(),
	}, nil
}

func (g *gaode) Platform() consts.PlatformID {
	return consts.Gaode
}
