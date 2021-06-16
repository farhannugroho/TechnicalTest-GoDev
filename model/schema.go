package model

type ArticleRequest struct {
	ID                int
	Title             string
	Lead              string
	Content           string
	ImageCoverURL     string
	ImageCoverCaption *string
	Status            *string

	Tag1ID             int
	Tag1Name           string
	CategoryDetailID   int
	CategoryDetailName string
}

type Article struct {
	ID                int
	Title             string
	Lead              string
	Content           string
	ImageCoverURL     string
	ImageCoverCaption string
	Status            string

	TagsDetail     []*Tag
	CategoryDetail *Category
}

type Tag struct {
	ID   int
	Name string
}

type Category struct {
	ID   int
	Name string
}
