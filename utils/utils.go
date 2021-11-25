package utils

import (
	"GoBIMS/utils/errmsg"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func ReturnJSON(c *gin.Context, httpStatus int, code int, data ...interface{}) {
	c.JSON(httpStatus, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"user":    data,
	})

}

func RandString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
