package mysql

import (
	"apihut-server/models"
)

func CreateGreet(greet *models.Greet) error {
	return db.Create(&greet).Error
}

func GetGreetFromID(id uint) (*models.Greet, error) {
	var g models.Greet
	err := db.Where("id", id).First(&g).Error
	return &g, err
}

func GetGreetList() ([]*models.Greet, error) {
	var ls = make([]*models.Greet, 0)
	err := db.Find(&ls).Error
	return ls, err
}
