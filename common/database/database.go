package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"

	"gorm.io/driver/mysql"
)

var (
	DB *gorm.DB
)

// InitDB 初始化数据库连接
func InitDB() {
	viper.SetConfigFile("../common/config.ini")
	viper.ReadInConfig()

	database := viper.GetString("db.database")
	host := viper.GetString("db.host")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	port := viper.GetString("db.port")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var connectorErr error
	DB, connectorErr = gorm.Open(mysql.New(mysql.Config{
		DSN:                     dsn,
		DontSupportRenameColumn: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})
	if connectorErr != nil {
		log.Fatalf("Failed to connect to database: %s", fmt.Sprintf("%v", connectorErr))
	}

	// 设置连接池配置（可选）
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: ", fmt.Sprintf("%v", connectorErr.Error()))
	}

	// 设置连接池大小等配置（根据实际情况进行调整）
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}
