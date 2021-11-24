package utils

import (
	"GoBIMS/utils/errmsg"
	"math/rand"

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
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
