package model

import (
	"GoBIMS/utils/errmsg"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	PassWord string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("UserName = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE //200
}

// 新增用户
func CreatUser(user *User) (code int) {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}
