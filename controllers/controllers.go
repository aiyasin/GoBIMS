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

// 用户注册
func JoinUp(c *gin.Context) {
	var user model.User
	var code int
	_ = c.ShouldBindJSON(&user)
	code = model.CheckUser(user.UserName)
	if code == errmsg.SUCCSE {
		model.CreatUser(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"user":    user,
		"message": errmsg.GetErrMsg(code),
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
