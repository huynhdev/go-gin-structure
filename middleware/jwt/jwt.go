package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-structure/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
