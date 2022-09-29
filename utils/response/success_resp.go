package utils

import "github.com/gin-gonic/gin"

// SuccessMessage ...
func SuccessMessage(c *gin.Context, status int, msg string) *gin.Context {
	c.JSON(status, gin.H{
		"Code":    "0000",
		"Message": msg,
	})
	return c
}

// SuccessData ...
func SuccessData(c *gin.Context, status int, data interface{}) *gin.Context {
	c.JSON(status, gin.H{
		"Response_Code" : "0000",
		"Response_Data" : data,
	})

	return c
}
