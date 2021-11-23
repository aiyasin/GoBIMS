package controllers

import (
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

// 查询用户是否存在
func IsUserExist(c *gin.Context) {
	//
}

// 添加用户
func AddUser(c *gin.Context) {
	//
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
