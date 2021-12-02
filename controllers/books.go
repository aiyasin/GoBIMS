package controllers

import (
	"GoBIMS/model"
	"GoBIMS/utils"
	"GoBIMS/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取图书列表
func GetBookList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.CheckBookPage(pageSize, pageNum)
	if code == errmsg.SUCCESS {
		utils.ReturnJSON(c, http.StatusOK, code, total, data)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1, nil)
	}
}

// GetBookList 查询图书列表
func SearchBook(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	name := c.Query("name")
	categories := c.Query("categories")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.SearchBookInfo(name, categories, pageSize, pageNum)
	if code == errmsg.SUCCESS {
		utils.ReturnJSON(c, http.StatusOK, code, total, data)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1, nil)
	}
}

// AddBook 添加图书
func AddBook(c *gin.Context) {
	var data model.Book
	_ = c.ShouldBindJSON(&data)

	code := model.CreateBookInfo(&data)

	if code == errmsg.SUCCESS {
		utils.ReturnJSON(c, http.StatusOK, code, -1, data)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1, nil)
	}
}

// EditBook 编辑图书信息
func EditBook(c *gin.Context) {
	var data model.Book
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.EditBookInfo(id, &data)

	if code == errmsg.SUCCESS {
		utils.ReturnJSON(c, http.StatusOK, code, -1, nil)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1, nil)
	}
}

// DeleteBook 删除文章
func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteBookInfo(id)

	if code == errmsg.SUCCESS {
		utils.ReturnJSON(c, http.StatusOK, code, -1, nil)
	} else {
		utils.ReturnJSON(c, http.StatusUnprocessableEntity, code, -1, nil)
	}
}
