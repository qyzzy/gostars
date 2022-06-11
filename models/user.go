package models

import (
	"golang.org/x/crypto/bcrypt"
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
	var user User
	db.Select("username = ?", username).First(&user)
	if user.ID > 0 {
		return code.ErrorUsernameUsed
	}
	return code.SUCCESS
}

func GetUserByID(id int) (User, int) {
	var user User
	err := db.Table(userTableName()).Where("id = ?", id).Limit(1).Find(&user).Error
	if err != nil {
		return user, code.ERROR
	}
	return user, code.SUCCESS
}

func CheckLoginFront(username, password string) (User, int) {
	var user User
	var PasswordErr error

	err := db.Table(userTableName()).Where("username = ?").First(&user).Error
	if err != nil {
		return user, code.ERROR
	}

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if PasswordErr != nil {
		return user, code.ErrorPasswordWrong
	}

	return user, code.SUCCESS
}
