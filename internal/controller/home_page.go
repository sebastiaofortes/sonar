package controller

import "github.com/gin-gonic/gin"

func HomePage (ctx *gin.Context) {
	ctx.HTML(200, "home-page.html", gin.H{})
}