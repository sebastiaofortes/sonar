package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastiaofortes/sonar/internal/controller"
	"github.com/sebastiaofortes/sonar/internal/platform/dao"
	"github.com/sebastiaofortes/sonar/internal/routes"
	"github.com/sebastiaofortes/sonar/internal/service"
)

func main() {
	PostDao := dao.NewPostDao()
	PostService := service.NewDefaultPostService(PostDao)
	PostController := controller.NewPostHandler(&PostService)
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*.html")
	r.Static("/assets", "./web/assets")
	r = routes.ConfigurePostRoutes(r, PostController)
	r.Run()
}
