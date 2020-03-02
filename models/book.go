package models

type Book struct {
	ID int64 `json:"id"`
	Pages int64 `json:"pages"`
	Year int64 `json:"year"`
	Title string `json:"title"`
	Content string `json:"content"`
}