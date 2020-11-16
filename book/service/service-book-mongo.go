package service

import (
	"context"
	"time"

	"example.com/book"
	"example.com/models"
)

type BookMongoService struct {
	bookRepo       book.RepositoryMongo
	contextTimeout time.Duration
}

func NewBookMongoService(br book.RepositoryMongo, timeout time.Duration) book.ServiceMongo {
	return &BookMongoService{
		bookRepo:       br,
		contextTimeout: timeout,
	}
}

func (bs *BookMongoService) CreateBook(ctx context.Context, m models.BookMongo) error {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.CreateBook(ctx, m)
}

func (bs *BookMongoService) UpdateBook(ctx context.Context, id string, m models.BookMongo) error {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.UpdateBook(ctx, id, m)
}

func (bs *BookMongoService) DeleteBook(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.DeleteBook(ctx, id)
}

func (bs *BookMongoService) GetById(ctx context.Context, id string, res models.BookMongo) error {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.GetById(ctx, id, res)
}

func (bs *BookMongoService) GetAll(ctx context.Context) ([]models.BookMongo, error) {
	ctx, cancel := context.WithTimeout(ctx, bs.contextTimeout)
	defer cancel()

	return bs.bookRepo.GetAll(ctx)
}
