package ip_bank

import (
	"net"

	"github.com/apihutco/server/dao/mysql"
	"github.com/apihutco/server/models"
	"github.com/apihutco/server/utils/consts"
)

type local struct {
}

func LocalInit() IIPCtrl {
	return &local{}
}

func (l *local) GetIP(ip net.IP) (*models.IPBank, error) {
	return mysql.GetIPBank(ip)
}

func (l *local) Platform() consts.PlatformCode {
	return consts.RepoPlatform.Local
}
