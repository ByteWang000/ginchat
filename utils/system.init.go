package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// 解析app.yml配置文件
func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited ......")
	//fmt.Println("config app:", viper.Get("app"))
	//fmt.Println("config mysql:", viper.Get("mysql"))
}

func InitMySQL() {
	// 自定义日志模板，打印SQL语句
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 以1秒为阈值记录慢日志
			LogLevel:      logger.Info, // 设置日志记录级别为info
			Colorful:      true,        //记录日志时是否添加颜色
		})

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newlogger})
	fmt.Println("MySQL inited ......")
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println(user)
}
