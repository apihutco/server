package ip_bank

import (
	"net"

	"apihut-server/dao/mysql"
	"apihut-server/models"
	"apihut-server/utils/consts"
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
	return consts.Local
}
