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
	Username      string    `gorm:"type:varchar(20);not null" json:"username"`
	Password      string    `gorm:"type:varchar(500);not null" json:"password"`
	Role          int       `gorm:"default:2" json:"role"`
	Gender        int       `gorm:"type:int;not null" json:"gender"`
	Age           int       `gorm:"type:int;not null" json:"age"`
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

func userTableName() string {
	return "users"
}

func CreateUser(data *User) int {
	err := db.Table(userTableName()).Create(&data).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func CheckUser(username string) int {
	if username == "" {
		return code.ERROR
	}
	fmt.Println(username)
	var user User
	db.Table(userTableName()).Select("id").Where("username = ?", username).First(&user)
	fmt.Println(user)
	if user.ID > 0 {
		return code.ErrorUsernameUsed
	}
	return code.SUCCESS
}

func CheckUpUser(id int, username string) int {
	var user User
	db.Table(userTableName()).Select("id, username").Where("username = ?", username).First(&user)
	if user.ID == uint(id) {
		return code.SUCCESS
	}
	if user.ID > 0 {
		return code.ErrorUsernameUsed
	}
	return code.SUCCESS
}

func GetUserByID(id int) (User, int) {
	var user User
	err := db.Table(userTableName()).Where("id = ?", id).Limit(1).First(&user).Error
	if err != nil {
		return user, code.ERROR
	}
	return user, code.SUCCESS
}

func CheckLoginFront(username, password string) (User, int) {
	var user User
	//var PasswordErr error

	err := db.Table(userTableName()).Where("username = ?", username).First(&user).Error

	if err != nil {
		return user, code.ERROR
	}

	//PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//if PasswordErr != nil {
	//	return user, code.ErrorPasswordWrong
	//}

	return user, code.SUCCESS
}

func GetUsers(pageSize, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	err := db.Table(userTableName()).Select("id, username, role, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)

	if err != nil {
		return users, total
	}
	db.Model(&users).Count(&total)

	return users, total
}

func GetUsersByUsername(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Table(userTableName()).Select("id, username, role, created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}

	return users, total
}

func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Table(userTableName()).Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func DeleteUser(id int) int {
	var user User
	err := db.Table(userTableName()).Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}
