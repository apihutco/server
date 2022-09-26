package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strconv"
)

var ShareConf *AppConf

type AppConf struct {
	*Site   `mapstructure:"site"`
	*DB     `mapstructure:"db"`
	*Redis  `mapstructure:"redis"`
	*Logger `mapstructure:"logger"`
}

type Site struct {
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
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

func Init() error {
	viper.SetConfigFile("./conf/config.sample.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		_ = viper.Unmarshal(&ShareConf)
	})
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err = viper.Unmarshal(&ShareConf); err != nil {
		return err
	}
	return nil
}

func GetSitePort() string {
	return fmt.Sprintf(":%s", strconv.Itoa(ShareConf.Site.Port))
}
