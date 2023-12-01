package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath("./config")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println(viper.Get("postgresql"))            // map[port:3306 url:127.0.0.1]
	fmt.Println(viper.Get("postgresql.searchPath")) // 127.0.0.1

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v "+
		"sslmode=disable TimeZone=Asia/Shanghai search_path=%v",
		viper.Get("postgresql.host"),
		viper.Get("postgresql.port"),
		viper.Get("postgresql.user"),
		viper.Get("postgresql.password"),
		viper.Get("postgresql.dbname"),
		viper.Get("postgresql.searchPath"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
	})
	if err != nil {
		fmt.Println(err)
	}
	//打印sql
	DB = DB.Debug()
	//设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.SetMaxOpenConns(10) // 最大连接数
	sqlDB.SetMaxIdleConns(5)  // 空闲连接数
}
