package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"ginchat/models"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err :=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))
}

func InitMySQL() {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println(user)
}
