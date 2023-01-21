package service

import (
	"github.com/sebastiaofortes/sonar/internal/domain"
	"github.com/sebastiaofortes/sonar/internal/platform/dao"
)

type DafeaulPostService struct {
	r domain.PostRepository
}

func (d *DafeaulPostService) GetPost(id int64) (domain.Post, error) {
	return d.r.GetPost(id)
}

func (d *DafeaulPostService) GetCategories() ([]domain.Categories, error) {
	return d.r.GetCategories()
}

func (d *DafeaulPostService) GetPostsByCategory(id int64) ([]domain.Post, error) {
	return d.r.GetPostsByCategory(id)
}

func NewDefaultPostService(domain.PostRepository) DafeaulPostService {
	d := dao.NewPostDao()
	return DafeaulPostService{r: d}
}
