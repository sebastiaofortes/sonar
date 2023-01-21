package domain

type Post struct {
	CategorieID int64
	Id          int64
	Title       string
	Text        string
	Data        string
}

type Categories struct {
	Id   int64
	Name string
}

type PostDao interface {
	GetPost(id string) (Post, error)
	GetCategories() ([]Categories, error)
}
