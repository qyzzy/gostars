package service

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gostars/global"
	"gostars/models"
	"gostars/utils/code"
)

type UserService struct {
}

func (userService *UserService) Register(data *models.User) int {
	username := data.Username
	// check username
	errCode := userService.CheckUser(username)
	if errCode != code.SUCCESS {
		return errCode
	}
	// generate uuid
	data.UUID = uuid.NewV4()
	err := global.GDb.Table(models.UserTableName()).Create(&data).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (userService *UserService) CheckUser(username string) int {
	if username == "" {
		return code.ERROR
	}
	var user *models.User
	err := global.GDb.Table(models.UserTableName()).
		Select("id").Where("username = ?", username).First(&user)
	if err != nil {
		return code.ERROR
	}
	if user.ID > 0 {
		return code.ErrorUsernameUsed
	}
	return code.SUCCESS
}

// admin right
func (userService *UserService) CreateUser(data *models.User) int {
	data.UUID = uuid.NewV4()
	err := global.GDb.Table(models.UserTableName()).
		Create(&data).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (userService *UserService) GetUserByID(id int) (models.User, int) {
	var user models.User
	err := global.GDb.Table(models.UserTableName()).
		Where("id = ?", id).Limit(1).First(&user).Error
	if err != nil {
		return user, code.ERROR
	}
	return user, code.SUCCESS
}

func (userService *UserService) CheckLoginFront(username, password string) (models.User, int) {
	var user models.User
	var PasswordErr error

	err := global.GDb.Table(models.UserTableName()).
		Where("username = ?", username).First(&user).Error

	if err != nil {
		return user, code.ERROR
	}

	// check password
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if PasswordErr != nil {
		return user, code.ErrorPasswordWrong
	}

	return user, code.SUCCESS
}

func (userService *UserService) CheckLogin(username string, password string) (models.User, int) {
	var user models.User
	var PasswordErr error

	err := global.GDb.Table(models.UserTableName()).
		Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, code.SUCCESS
	}

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, code.ErrorUserNotExist
	}
	if PasswordErr != nil {
		return user, code.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return user, code.ErrorUserNoRight
	}
	return user, code.SUCCESS
}

func (userService *UserService) ChangePassword(data *models.User, newPassword string) int {
	var user *models.User
	err := global.GDb.Table(models.UserTableName()).
		Where("id = ?", data.ID).
		First(&user).Update("password", newPassword).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (userService *UserService) GetUsers(pageSize, pageNum int) ([]models.User, int64) {
	var users []models.User
	var total int64

	err := global.GDb.Table(models.UserTableName()).
		Select("id, username, role, created_at").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)

	if err != nil {
		return users, total
	}
	global.GDb.Model(&users).Count(&total)

	return users, total
}

func (userService *UserService) GetUsersByUsername(username string, pageSize int, pageNum int) ([]models.User, int64) {
	var users []models.User
	var total int64

	if username != "" {
		global.GDb.Table(models.UserTableName()).Select("id, uuid, username, role, created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		global.GDb.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}

	return users, total
}

func (userService *UserService) EditUser(id int, data *models.User) int {
	var user models.User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := global.GDb.Table(models.UserTableName()).
		Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (userService *UserService) DeleteUser(id int) int {
	var user models.User
	err := global.GDb.Table(models.UserTableName()).
		Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}
