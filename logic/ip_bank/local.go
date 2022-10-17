package ip_bank

import (
	"apihut-server/dao/mysql"
	"apihut-server/logic/consts"
	"apihut-server/models"
	"net"
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
