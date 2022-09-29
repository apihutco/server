package mysql

import (
	"apihut-server/models"
	"net"
)

func CreateIPBank(ip *models.IPBank) error {
	return db.Create(&ip).Error
}

func GetIPBank(ip net.IP) (*models.IPBank, error) {
	info := new(models.IPBank)
	err := db.Where("ip", ip.String()).First(&info).Error
	return info, err
}