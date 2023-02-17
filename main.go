package main

import (
	"ginchat/router"
	"ginchat/utils"
	// "gorm.io/gorm/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置文件
	utils.InitConfig()
	// 初始化数据库
	utils.InitMySQL()

	r := router.Router()

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func _json(c *gin.Context) {
	// 1.json响应结构体
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		PassWord string `json:"-"` // 忽略转换为json，不显示
	}
	//user := UserInfo{"三体", 23, "123456"}
	// 2.json响应map
	userMap := map[string]interface{}{
		"username": "四体",
		"age":      4,
	}
	c.JSON(200, userMap)
	// 3.直接响应json
	c.JSON(200, gin.H{"username": "五体", "age": 23})
}

func _string(c *gin.Context) {
	c.String(200, "你好")
}
