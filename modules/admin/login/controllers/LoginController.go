package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackjewl/squirrelBlog/modules/admin/login/entities"
	"github.com/jackjewl/squirrelBlog/modules/admin/login/model"
	"github.com/jackjewl/squirrelBlog/utils"
)

type LoginController struct {
	Request  interface{}
	Response interface{}
}

// Prepare 初始化controller
func (this *LoginController) Prepare(c *gin.Context) {

}

//登录
func (this *LoginController) Login(c *gin.Context) {

	//检查token
	s := c.Request.Header.Get("token")
	if s != "" {
		r, _ := utils.ParseFromToken(s)

		//如果toekn有且合法，直接不用登录
		if r.RuleType == utils.Admin {
			return
		}
	}
	this.Prepare(c)

	var a entities.Admin

	if r, _ := model.AdminLogin(a); r {

		token, _ := utils.GenerateToken(1, a.Email)

		//设置token
		c.Request.Response.Header.Add("token", token)

		c.Status(http.StatusHTTPVersionNotSupported)
		c.Abort()
		return
	}
}
