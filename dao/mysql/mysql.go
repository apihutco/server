package mysql

import (
	"apihut-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {

	switch config.ShareConf.DB.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.ShareConf.DB.SQLite.Name), &gorm.Config{})
	default:
		db, err = gorm.Open(mysql.Open(""), &gorm.Config{})
	}

	err = db.AutoMigrate()
	return err
}
