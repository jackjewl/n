package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jackjewl/squirrelBlog/utils"

	"github.com/gin-gonic/gin"
)

// TokenValid 在每次登录前拦截验证token
func TokenValid(check bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		if !check {
			c.Next()
		} else {
			token := c.Request.Header.Get("token")

			if token == "" {
				c.Abort()
				c.Status(http.StatusUnauthorized)
				c.Error(errors.New("需要登录"))
				return
			}

			if strings.HasPrefix(c.Request.URL.Path, "/admin") {

			}

			rule, err := utils.ParseFromToken(token)
			if err != nil {
				c.Abort()
				c.Status(http.StatusUnauthorized)
				c.Error(errors.New("需要登录"))
				return
			}

			if rule.Account == "" {
				c.Abort()
				c.Status(http.StatusUnauthorized)
				c.Error(errors.New("需要登录"))
				return
			}

			c.Next()
		}

	}
}
