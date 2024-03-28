/*
 * @Author: yujiajie
 * @Date: 2024-03-13 11:43:40
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 14:33:31
 * @FilePath: /stage/sdk/config/app.go
 * @Description:
 */
package config

import (
	"stage/sdk/core"

	cfg "github.com/bird-coder/manyo/config"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Server    *cfg.HttpConfig             `yaml:"server"`
	Logger    *cfg.LoggerConfig           `yaml:"logger"`
	Databases map[string]*cfg.MysqlConfig `yaml:"databases" mapstructure:"databases"`
	Cache     *cfg.CacheConfig            `yaml:"cache"`
	Locker    *cfg.LockConfig             `yaml:"locker"`
}

func (app *AppConfig) LoadConfig(configFile string) (err error) {
	viper.SetConfigFile(configFile)
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&app); err != nil {
		return
	}
	core.App.SetConfig("http", app.Server)
	core.App.SetConfig("logger", app.Logger)
	core.App.SetConfig("databases", app.Databases)
	core.App.SetConfig("cache", app.Cache)
	core.App.SetConfig("locker", app.Locker)
	return
}
