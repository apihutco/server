package db

import (
	"sync"

	"github.com/apihutco/server/models"
	"gorm.io/gorm"
)

type IGreet interface {
	Create(greet *models.Greet) error
	GetByID(id uint) (*models.Greet, error)
	List() ([]*models.Greet, error)
	Update(greet *models.Greet) error
}

type greetCtrl struct {
	db *gorm.DB
}

var shareGreetCtrl *greetCtrl
var greetCtrlOnce sync.Once

func (data *Database) Greet() IGreet {
	greetCtrlOnce.Do(func() {
		shareGreetCtrl = &greetCtrl{
			db: data.db,
		}
	})
	return shareGreetCtrl
}

func (g *greetCtrl) Create(greet *models.Greet) error {
	return g.db.Create(&greet).Error
}

func (g *greetCtrl) GetByID(id uint) (*models.Greet, error) {
	var greet models.Greet
	err := g.db.Where("id", id).First(&greet).Error
	return &greet, err
}

func (g *greetCtrl) List() ([]*models.Greet, error) {
	var ls = make([]*models.Greet, 0)
	err := g.db.Find(&ls).Error
	return ls, err
}

func (g *greetCtrl) Update(greet *models.Greet) error {
	return g.db.Where("id", greet.ID).Updates(&greet).Error
}
