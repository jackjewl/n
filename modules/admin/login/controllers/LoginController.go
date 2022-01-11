package controllers

import "github.com/gin-gonic/gin"

type LoginController struct {
	Request  interface{}
	Response interface{}
}

// Prepare 初始化controller
func (this *LoginController) Prepare(ctx *gin.Context) {

}