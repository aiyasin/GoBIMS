package controllers

import (
	"GoBIMS/model"
	"GoBIMS/utils/errmsg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户登录
func Login(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")

	c.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", username, password, types))
}

// 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	var validCode int
	_ = c.ShouldBindJSON(&user)
	validCode = model.CheckUser(user.UserName)
	// msg, validCode := validator.Validate(&user)
	if validCode == errmsg.SUCCSE {
		model.CreatUser(&user)
	}
	if validCode == errmsg.ERROR_USERNAME_USED {
		validCode = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  validCode,
		"user":    user,
		"message": errmsg.GetErrMsg(validCode),
	})
}

// 查询用户
func GetUser(c *gin.Context) {
	//
}

// 编辑用户
func EditUser(c *gin.Context) {
	//
}

// 删除用户
func DeleteUser(c *gin.Context) {
	//
}
