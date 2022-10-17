package mysql

import (
	"apihut-server/config"
	"apihut-server/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {

	switch config.Share.DB.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.Share.DB.SQLite.Name), &gorm.Config{})
	default:
		db, err = gorm.Open(mysql.Open(""), &gorm.Config{})
	}

	err = db.AutoMigrate(
		&models.IPBank{},
		&models.Greet{},
	)
	return err
}
