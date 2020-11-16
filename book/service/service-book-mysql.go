package service

import (
	"context"
	"time"

	"example.com/book"
	"example.com/models"
)

type bookService struct {
	bookRepo       book.Repository
	contextTimeout time.Duration
}

func NewBookService(br book.Repository, timeout time.Duration) book.Service {
	return &bookService{
		bookRepo:       br,
		contextTimeout: timeout,
	}
}

func (bs *bookService) CreateBook(ctx context.Context, m models.Book) error {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.CreateBook(ctx, m)
}

func (bs *bookService) UpdateBook(ctx context.Context, id int64, m models.Book) error {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.UpdateBook(ctx, id, m)
}

func (bs *bookService) DeleteBook(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.DeleteBook(ctx, id)
}

func (bs *bookService) GetById(ctx context.Context, id int64, res models.Book) error {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.GetById(ctx, id, res)
}

func (bs *bookService) GetAll(ctx context.Context) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.GetAll(ctx)
}
