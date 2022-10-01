package middleware

import (
	"belajar/efishery/auth"
	"belajar/efishery/configs"
	utils "belajar/efishery/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware(config *configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		_,err := auth.TokenValid(c.Request,"",config)
		if err != nil {
				utils.ErrorMessage(c,http.StatusUnauthorized,err.Error(),401)
			c.Abort()
			return
		}
		c.Next()
	}
}

