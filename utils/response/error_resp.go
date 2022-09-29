package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// ErrorMessage ...
func ErrorMessage(c *gin.Context, status int, msg string, code int) *gin.Context {
	c.JSON(status, gin.H{
		"Code": code,
		"Message": msg,
	})
	return c
}


// func to create error from string
func NewErr(errs string) error {
	return errors.New(errs)
}