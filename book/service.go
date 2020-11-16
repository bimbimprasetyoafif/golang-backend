package book

import (
	"context"

	"example.com/models"
)

type Service interface {
	CreateBook(ctx context.Context, m models.Book) error
	DeleteBook(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64, m models.Book) error
	GetAll(ctx context.Context) ([]models.Book, error)
	UpdateBook(ctx context.Context, id int64, m models.Book) error
}

type ServiceMongo interface {
	CreateBook(ctx context.Context, m models.BookMongo) error
	DeleteBook(ctx context.Context, id string) error
	GetById(ctx context.Context, id string, m models.BookMongo) error
	GetAll(ctx context.Context) ([]models.BookMongo, error)
	UpdateBook(ctx context.Context, id string, m models.BookMongo) error
}
