package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebastiaofortes/sonar/internal/service"
)

type PostHandler struct {
	postService service.PostService
}

func (h *PostHandler) GetPost(templateName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, err := h.postService.GetCategories()
		if err != nil {
			ctx.HTML(404, templateName, gin.H{})
		}

		id := ctx.Query("id")
		if id == "" {
			ctx.HTML(200, templateName, gin.H{
				"categories": c,
			})
			return
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.HTML(400, templateName, gin.H{
				"categories": c,
			})
			return
		}
		p, err := h.postService.GetPost(int64(idInt))
		if err != nil {
			ctx.HTML(404, templateName, gin.H{
				"categories": c,
			})
		}

		ctx.HTML(200, templateName, gin.H{
			"text":       p.Text,
			"title":      p.Title,
			"categories": c,
		})
	}
}

func (h *PostHandler) ListCategories(templateName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, err := h.postService.GetCategories()
		if err != nil {
			ctx.HTML(404, templateName, gin.H{})
		}

		id := ctx.Query("id")
		if id == "" {
			ctx.HTML(200, templateName, gin.H{
				"categories": c,
			})
			return
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.HTML(400, templateName, gin.H{
				"categories": c,
			})
			return
		}
		p, err := h.postService.GetPostsByCategory(int64(idInt))
		if err != nil {
			ctx.HTML(404, templateName, gin.H{
				"categories": c,
			})
		}

		ctx.HTML(200, templateName, gin.H{
			"posts":      p,
			"categories": c,
		})
	}
}

func NewPostHandler(s service.PostService) PostHandler {
	return PostHandler{postService: s}
}
