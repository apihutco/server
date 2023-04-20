package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var Conf *AppConf

type AppConf struct {
	*Site   `mapstructure:"site"`
	*DB     `mapstructure:"db"`
	*Redis  `mapstructure:"redis"`
	*Logger `mapstructure:"logger"`
	*Open   `mapstructure:"open"`
	*File   `mapstructure:"file"`
	*Bleve  `mapstructure:"bleve"`
}

type Site struct {
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`
	DocsUrl string `mapstructure:"docs_url"`
}

type DB struct {
	Driver  string `mapstructure:"driver"`
	*MySQL  `mapstructure:"mysql"`
	*SQLite `mapstructure:"sqlite"`
}

type MySQL struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
}

type SQLite struct {
	Name string `mapstructure:"name"`
}

type Redis struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Logger struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init(f string) error {
	viper.SetConfigFile(f)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err = viper.Unmarshal(&Conf); err != nil {
		return err
	}

	// 创建目录
	err = os.MkdirAll(Conf.File.Avatar, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (conf *AppConf) GetSitePort() string {
	return fmt.Sprintf(":%s", strconv.Itoa(conf.Site.Port))
}
