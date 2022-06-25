package models

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	UUID          uuid.UUID `json:"uuid"`
	Avatar        string    `json:"avatar"`
	Username      string    `gorm:"type:varchar(20);not null" json:"username"`
	Password      string    `gorm:"type:varchar(500);not null" json:"password"`
	AuthorityID   string    `gorm:"not null;size:90" json:"authorityid"`
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

func UserTableName() string {
	return "users"
}

// BeforeCreate encrypt
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// Build password
func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}
