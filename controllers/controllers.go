package controllers

import (
	"GoBIMS/model"
	"GoBIMS/utils"
	"GoBIMS/utils/errmsg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
	code = model.CheckUser(user)
	if code == errmsg.SUCCSE {
		if len(user.UserName) == 0 {
			user.UserName = utils.RandString(10)
		}
		model.CreatUser(&user)
		utils.ReturnJSON(c, http.StatusOK, code, user)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code)
	}
}

// 查询用户
func GetUser(c *gin.Context) {
	pageSize := cast.ToInt(c.Query("pagesize"))
	pageNum := cast.ToInt(c.Query("pagenum"))
	username := c.Query("user_name")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.CheckUserPage(username, pageSize, pageNum)
	code := errmsg.SUCCSE
	pagenumtotal := fmt.Sprintf("第%d页/共%d个用户", pageNum, total)
	utils.ReturnJSON(c, http.StatusOK, code, data, pagenumtotal)
}

// 编辑用户
func EditUser(c *gin.Context) {
	//
}

// 删除用户
func DeleteUser(c *gin.Context) {
	//
}
