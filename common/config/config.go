package config

import (
	"github.com/spf13/viper"
)

var Config = make(map[string]string)

func InitConfig(env string) map[string]string {
	if env == "local" {
		viper.SetConfigFile("../common/local_config.ini")
	} else if env == "test" {
		viper.SetConfigFile("../common/test_config.ini")
	} else if env == "prod" {
		viper.SetConfigFile("../common/prod_config.ini")
	}
	viper.ReadInConfig()
	Config["app_name"] = viper.GetString("app.name")
	Config["app_domain"] = viper.GetString("app.domain")

	Config["db_database"] = viper.GetString("db.database")
	Config["db_host"] = viper.GetString("db.host")
	Config["db_username"] = viper.GetString("db.username")
	Config["db_password"] = viper.GetString("db.password")
	Config["db_port"] = viper.GetString("db.port")

	Config["redis_host"] = viper.GetString("redis.host")
	Config["redis_port"] = viper.GetString("redis.port")
	Config["redis_password"] = viper.GetString("redis.password")

	return Config
}
