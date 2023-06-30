package geo_bank

import (
	"github.com/apihutco/server/dao/db"
	"github.com/apihutco/server/models"
)

var _ IGeoCtrl = &GeoBase{}

type GeoBase struct{}

func (g *GeoBase) GetInfo(location string, adm string) (*models.GeoBank, error) {
	panic("implement me")
}

func (g *GeoBase) SaveInfo(info *models.GeoBank) error {
	return db.Ctrl().Geo().Create(info)
}
