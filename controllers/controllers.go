package controllers

import (
	"GoBIMS/model"
	"GoBIMS/utils"
	"GoBIMS/utils/errmsg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Login 用户登录
func Login(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")

	c.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", username, password, types))
}

// JoinUp 用户注册
func JoinUp(c *gin.Context) {
	var user model.User
	var code int
	_ = c.ShouldBindJSON(&user)
	code = model.CheckUser(user)
	if code == errmsg.SUCCESS {
		model.CreatUser(&user)
		utils.ReturnJSON(c, http.StatusOK, code, user)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code)
	}
}

// GetUser 查询用户列表
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
	data, totalnum := model.CheckUserPage(username, pageSize, pageNum)
	code := errmsg.SUCCESS
	pagenumtotal := fmt.Sprintf("第%d页|共%d页", pageNum, totalnum)
	utils.ReturnJSON(c, http.StatusOK, code, data, pagenumtotal)
}

// EditUser 编辑用户

func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUpUser(id, data.UserName)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	utils.ReturnJSON(c, http.StatusOK, code)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	utils.ReturnJSON(c, http.StatusOK, code)
}
