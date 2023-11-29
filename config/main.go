package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// 用viper读取yaml文件，感觉比yaml.v3好用多了
func main() {
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
	fmt.Println(viper.Get("postgresql.searchPath"))
	sprintf := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v "+
		"sslmode=disable TimeZone=Asia/Shanghai search_path=%v",
		viper.Get("postgresql.host"),
		viper.Get("postgresql.port"),
		viper.Get("postgresql.user"),
		viper.Get("postgresql.password"),
		viper.Get("postgresql.dbname"),
		viper.Get("postgresql.searchPath"),
	)
	fmt.Printf("%v", sprintf)
}
