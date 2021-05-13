package model

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRepository interface {
	CreateTable()
	InsertBook(title string, author string) int64
	GetBooks() []Book
	GetBook(id int64) Book
}
