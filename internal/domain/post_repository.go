package domain

type PostRepository interface{
	GetPost(id int64) (Post, error)
	GetCategories() ([]Categories, error)
	GetPostsByCategory(id int64) ([]Post, error)
}