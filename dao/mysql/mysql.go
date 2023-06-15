package mysql

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"apihut-server/config"
	"apihut-server/logger"
	"apihut-server/models"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {

	switch config.Conf.DB.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.Conf.DB.SQLite.Name), &gorm.Config{})
	default:
		db, err = gorm.Open(mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.Conf.DB.MySQL.User,
				config.Conf.DB.MySQL.Password,
				config.Conf.DB.MySQL.Host,
				config.Conf.DB.MySQL.DBName,
			),
		), &gorm.Config{})
	}

	err = db.AutoMigrate(
		&models.IPBank{},
		&models.Greet{},
		&models.GeoBank{},
	)
	if err != nil {
		return err
	}

	// 暂时屏蔽,手动初始化SQL
	//err = loadSQLFile()

	return err
}

// 载入所有初始化sql文件
func loadSQLFile() error {
	fileList, err := os.ReadDir(config.Conf.Bleve.SetupPath)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	for _, entry := range fileList {
		wg.Add(1)
		go func(entry os.DirEntry) {
			file, err := os.Open(path.Join(config.Conf.Bleve.SetupPath, entry.Name()))
			defer func() {
				_ = file.Close()
				wg.Done()
			}()
			if err != nil {
				logger.L().Error("Setup open file", zap.Error(err))
				return
			}
			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				logger.L().Error("Setup read all", zap.Error(err))
				return
			}

			logger.L().Debug("Setup SQL", zap.ByteString("content", bytes))

			err = Exec(string(bytes))
			if err != nil {
				logger.L().Error("Setup sql exec", zap.Error(err))
				return
			}
		}(entry)
	}

	wg.Wait()

	return nil
}
