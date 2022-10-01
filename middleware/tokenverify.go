package middleware

import (
	"belajar/efishery/auth"
	utils "belajar/efishery/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_,err := auth.TokenValid(c.Request,"")
		if err != nil {
				utils.ErrorMessage(c,http.StatusUnauthorized,err.Error(),401)
			c.Abort()
			return
		}
		c.Next()
	}
}

