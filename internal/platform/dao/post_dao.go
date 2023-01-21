package dao

import (
	"errors"

	"github.com/sebastiaofortes/sonar/internal/domain"
)

var mBD []domain.Post

var cBD []domain.Categories

type PostDao struct {
}

func (d PostDao) GetPost(id int64) (domain.Post, error) {
	for _, v := range mBD {
		if v.Id == id {
			return v, nil
		}
	}
	return domain.Post{}, errors.New("[dao GetPost] Nenhum post encontrado")
}

func (d PostDao) GetCategories() ([]domain.Categories, error) {
	return cBD, nil
}

func (d PostDao) GetPostsByCategory(id int64) ([]domain.Post, error) {
	res := []domain.Post{}
	for _, v := range mBD {
		if v.CategorieID == id {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return res, errors.New("[dao GetPostByCategory] Nenhum post encontrado")
	}
	return res, nil
}

func NewPostDao() PostDao {
	cBD = []domain.Categories{
		{Id: 12, Name: "Noticias"},
		{Id: 22, Name: "Recados"},
		{Id: 33, Name: "Avisos"},
	}

	mBD = []domain.Post{
		{Id: 1, CategorieID: 12, Title: "Meu 1 título", Text: "Meu texto"},
		{Id: 2, CategorieID: 22, Title: "Meu 2 título", Text: "Meu texto"},
		{Id: 3, CategorieID: 33, Title: "Meu 3 título", Text: "Meu texto"},
	}
	return PostDao{}
}
