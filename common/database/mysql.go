package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	DB *gorm.DB
)

// InitDB 初始化数据库连接
func InitDB() {

	//database := projectConfig["db_database"]
	//host := projectConfig["db_host"]
	//username := projectConfig["db_username"]
	//password := projectConfig["db_password"]
	//port := projectConfig["db_port"]

	viper.SetConfigFile("./config.ini")
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
		log.Println("❌数据库链接失败", connectorErr)
	}

	// 设置连接池配置（可选）
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("❌设置数据库链接失败", err)
	}
	log.Println("✅ 数据库链接成功")

	// 设置连接池大小等配置（根据实际情况进行调整）
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

}
