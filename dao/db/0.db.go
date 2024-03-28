package db

import (
	"fmt"

	"github.com/apihutco/server/config"
	"github.com/apihutco/server/models"
	//"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
	//mongo *mongo.Client
}

var sharedDatabase *Database

func Init() (err error) {
	sharedDatabase = new(Database)
	switch config.Conf.DB.Driver {
	case "sqlite":
		sharedDatabase.db, err = gorm.Open(sqlite.Open(config.Conf.DB.SQLite.Name), &gorm.Config{})
	default:
		sharedDatabase.db, err = gorm.Open(mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.Conf.DB.MySQL.User,
				config.Conf.DB.MySQL.Password,
				config.Conf.DB.MySQL.Host,
				config.Conf.DB.MySQL.DBName,
			),
		), &gorm.Config{})
	}

	err = sharedDatabase.db.AutoMigrate(
		&models.IPBank{},
		&models.Greet{},
		&models.GeoBank{},
	)
	if err != nil {
		return err
	}

	return err
}

func Ctrl() *Database {
	if sharedDatabase == nil {
		if err := Init(); err != nil {
			panic(err)
		}
	}
	return sharedDatabase
}
