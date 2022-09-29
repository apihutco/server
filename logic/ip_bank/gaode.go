package ip_bank

import (
	"apihut-server/logic/consts"
	"apihut-server/models"
	"net"
)

type gaode struct {
	key string
}

func GaodeInit(key string) IIPCtrl {
	return &gaode{key: key}
}

func (g *gaode) GetIP(ip net.IP) (*models.IPBank, error) {
	// TODO implement me
	panic("implement me")
}

func (g *gaode) Platform() consts.PlatformID {
	return consts.Gaode
}
