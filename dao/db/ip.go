package db

import (
	"net"
	"sync"

	"github.com/apihutco/server/models"
	"gorm.io/gorm"
)

type IIP interface {
	Create(ip *models.IPBank) error
	Get(ip net.IP) (*models.IPBank, error)
}

type ipCtrl struct {
	db *gorm.DB
}

var shareIPCtrl *ipCtrl
var ipCtrlOnce sync.Once

func (data *Database) IP() IIP {
	ipCtrlOnce.Do(func() {
		shareIPCtrl = &ipCtrl{
			db: data.db,
		}
	})
	return shareIPCtrl
}

func (i *ipCtrl) Create(ip *models.IPBank) error {
	return i.db.Create(&ip).Error
}

func (i *ipCtrl) Get(ip net.IP) (*models.IPBank, error) {
	info := new(models.IPBank)
	err := i.db.Where("ip", ip.String()).First(&info).Error
	return info, err
}
