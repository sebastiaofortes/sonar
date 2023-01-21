package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastiaofortes/sonar/internal/controller"
)

func ConfigurePostRoutes(r *gin.Engine, c controller.PostHandler) *gin.Engine {
	r.GET("/", controller.HomePage)
	r.GET("/mouse/categories", c.ListCategories("mouse-categories.html"))
	r.GET("mouse/post", c.GetPost("mouse-post.html"))
	r.GET("/teclado/categories", c.ListCategories("teclado-categories.html"))
	r.GET("/teclado/post", c.GetPost("teclado-post.html"))
	return r
}
