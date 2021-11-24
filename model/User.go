package model

import (
	"GoBIMS/utils/errmsg"
	"log"
	"math"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null;columns:user_name" json:"user_name" validate:"required,min=4,max=12" label:"用户名"`
	PassWord string `gorm:"type:varchar(500);not null columns:pass_word" json:"pass_word" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;default:2;columns:role" json:"role" validate:"required" label:"角色码"`
}

// CheckUser 查询用户是否存在
func CheckUser(user User) (code int) {
	db.Select("id").Where("user_name = ?", user.UserName).First(&user)
	if len(user.PassWord) < 6 {
		return errmsg.ErrorPasswordLessThan6
	}
	if user.ID > 0 {
		return errmsg.ErrorUsernameUsed //1001
	}
	return errmsg.SUCCESS //200
}

// CreatUser 新增用户
func CreatUser(user *User) (code int) {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// CheckUserPage 查询用户列表
func CheckUserPage(username string, pageSize int, pageNum int) ([]User, int64) {
	var user []User
	var total int64
	if username != "" {
		db.Select("id,user_name,role,created_at").Where(
			"user_name LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user)
		db.Model(&user).Where(
			"user_name LIKE ?", username+"%",
		).Count(&total)
		totalnum := int64(math.Ceil(float64(total) / float64(pageSize)))
		return user, totalnum
	}
	db.Select("id,user_name,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user)
	db.Model(&user).Count(&total)

	if err != nil {
		return user, 0
	}
	totalnum := int64(math.Ceil(float64(total) / float64(pageSize)))
	return user, totalnum
}

// BeforeCreate 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.PassWord = ScryptPassWord(u.PassWord)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.PassWord = ScryptPassWord(u.PassWord)
	return nil
}

// ScryptPassWord 密码加密
func ScryptPassWord(password string) string {
	const cost = 10
	HashPassWord, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPassWord)
}
