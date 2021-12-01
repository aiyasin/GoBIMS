package controllers

import (
	"GoBIMS/model"
	"GoBIMS/utils"
	"GoBIMS/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// JoinUp 用户注册
func JoinUp(c *gin.Context) {
	var user model.User
	var code int
	_ = c.ShouldBindJSON(&user)
	code = model.CheckUser(user)
	if code == errmsg.SUCCESS {
		model.CreatUser(&user)
		utils.ReturnJSON(c, http.StatusOK, code, -1, user)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1)
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
	data, code, total := model.CheckUserPage(username, pageSize, pageNum)
	if code == errmsg.SUCCESS {
		utils.ReturnJSON(c, http.StatusOK, code, total, data)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1)
	}
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUpUser(id, data.UserName)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
		utils.ReturnJSON(c, http.StatusOK, code, -1)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1)
	}
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	if code == errmsg.SUCCESS {
		utils.ReturnJSON(c, http.StatusOK, code, -1)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1)
	}
}
