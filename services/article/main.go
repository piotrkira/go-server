// Package article is an example of simple service
package article

// Article struct
type Article struct {
	Author  string
	Title   string
	Content string
}

// GetLatestArticle returns latest Article
func GetLatestArticle() Article {
	var a Article
	a.Author = "Piotr Kira"
	a.Title = "Title"
	a.Content = "It's an example content of latest post"
	return a
}
