package geo_bank

import (
	"github.com/apihutco/server/dao/db"
	"github.com/apihutco/server/models"
)

type Local struct {
	GeoBase
}

func (l *Local) GetInfo(location string, adm string) (*models.GeoBank, error) {
	return db.Ctrl().Geo().Get(location, adm)
}

func (l *Local) SaveInfo(info *models.GeoBank) error {
	return nil
}

func NewLocal() IGeoCtrl {
	return &Local{}
}
