package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	DB *gorm.DB
)

// InitDB 初始化数据库连接
func InitDB(projectConfig map[string]string) {

	database := projectConfig["db_database"]
	host := projectConfig["db_host"]
	username := projectConfig["db_username"]
	password := projectConfig["db_password"]
	port := projectConfig["db_port"]

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
