package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model    //匿名字段，UserBasic可以访问内部所有属性
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identidy      string
	ClentIp       string
	ClentPort     string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"` //使用了"gorm:"column:login_out_time" json:"login_out_time"`标签，代表在数据库中的列名为"login_out_time"，在序列化/反序列化 JSON 数据时的字段名为 "login_out_time".
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
