package models

type Post struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Body  string `db:"body" json:"body"`
}

func GetPost() Post {
	return Post{}
}

func GetPosts() []Post {
	return []Post{}
}