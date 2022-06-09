package models

import (
	"fmt"
	"gorm.io/gorm"
	"gostars/utils/code"
	"time"
)

type User struct {
	gorm.Model
	Avatar        string    `json:"avatar"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Role          int       `gorm:"default:2" json:"role"`
	Gender        int       `json:"gender"`
	Age           int       `json:"age"`
	Birthday      time.Time `json:"birthday"`
	Email         string    `json:"email"`
	City          string    `json:"city"`
	LastLoginTime time.Time `json:"lastlogintime"`
	Status        int       `json:"status"`
	Mobile        string    `json:"mobile"`
	LastIP        string    `json:"lastip"`
	IPSource      string    `json:"ipsource"`
	Browser       string    `json:"browser"`
}

func tableName() string {
	return "users"
}

func CreateUser(data *User) int {
	err := db.Table(tableName()).Create(&data).Error
	if err != nil {
		fmt.Println("Create user failed")
		panic(err)
	}
	return code.SUCCESS
}

func CheckUser(username string) int {
	return code.SUCCESS
}

func GetUser(id int) (User, int) {
	var user User
	err := db.Table(tableName()).Where("id = ?", id).Limit(1).Find(&user).Error
	if err != nil {
		return user, code.ERROR
	}
	return user, code.SUCCESS
}
