package config

import (
	"github.com/colinrs/pkgx/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	// Conf ...
	Conf *Config
)

type Config struct {
	Mysql  MySQLConfig `json:"mysql"`
	Kafka  KafkaConfig `json:"kafka"`
	Server ServerConfig `json:"server"`
}

// MySQLConfig ...
type MySQLConfig struct {
	Name            string `json:"name"`
	Addr            string `json:"addr"`
	DB              string `json:"db"`
	UserName        string `json:"username"`
	Password        string `json:"password"`
	MaxIdleConn     int    `json:"max_idel_conn"`
	MaxOpenConn     int    `json:"max_open_conn"`
	ConnMaxLifeTime int    `json:"conn_max_lifetime"`
}

type KafkaConfig struct {
	Addr string `json:"addr"`
}

type ServerConfig struct {
	Addr string `json:"addr"`
	Port int `json:"port"`
}


func Init(configPath string) error {
	if err := initConfig(configPath); err!=nil {
		return err
	}
	return nil
}

func initConfig(configPath string) error {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	}
	viper.SetConfigType("json")
	// 导入环境变量
	//viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	err := viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	logger.Info("config:(%#v)", Conf)
	//// 应该是配置文件热更新
	//watchConfig()
	return nil
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func( e fsnotify.Event) {
		logger.Info("Config file changed: %s", e.Name)
	})
}

