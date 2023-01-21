package service

import "github.com/sebastiaofortes/sonar/internal/domain"

type PostService interface {
	GetPost(id int64) (domain.Post, error)
	GetCategories() ([]domain.Categories, error)
	GetPostsByCategory(id int64) ([]domain.Post, error)
}
