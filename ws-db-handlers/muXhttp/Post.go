package muXhttp

type Post struct {
	Id int
	Title string
	Content string
}

func NewPost(id int, title, content string) *Post {
	return &Post{id, title, content}
}