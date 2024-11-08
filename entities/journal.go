package entities

type Journal struct {
	ID         uint
	Title      string
	Content    string
	AuthorID   uint
	CategoryID uint
	File       string
	Author     Author
	Category   Category
}
