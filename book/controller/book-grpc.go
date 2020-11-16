package controller

import (
	"context"
	"fmt"

	"example.com/book"
	"example.com/models"
	pb "example.com/pb"
)

type grpcServer struct {
	pb.BookGrpcServer
	book.ServiceMongo
}

func NewGrpc(bookService book.ServiceMongo) pb.BookGrpcServer {
	return &grpcServer{
		ServiceMongo: bookService,
	}
}

func (g *grpcServer) GetAllBook(ctx context.Context, message *pb.Message) (*pb.Books, error) {
	res, err := g.ServiceMongo.GetAll(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var allBooks []*pb.Book
	for _, v := range res {
		allBooks = append(allBooks, &pb.Book{
			ID:      v.ID.Hex(),
			Pages:   int32(v.Pages),
			Year:    int32(v.Year),
			Title:   v.Title,
			Content: v.Content,
		})
	}

	return &pb.Books{
		AllBook: allBooks,
	}, nil
}

func (g *grpcServer) GetByIdBook(ctx context.Context, id *pb.BookId) (*pb.Book, error) {
	var res models.BookMongo
	if err := g.ServiceMongo.GetById(ctx, id.ID, res); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &pb.Book{
		ID:      res.ID.String(),
		Pages:   int32(res.Pages),
		Year:    int32(res.Year),
		Title:   res.Title,
		Content: res.Content,
	}, nil
}

func (g *grpcServer) DelByIdBook(ctx context.Context, id *pb.BookId) (*pb.Message, error) {
	if err := g.ServiceMongo.DeleteBook(ctx, id.ID); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.Message{
		Message: "Deleted",
		Error:   "",
	}, nil
}

func (g *grpcServer) CreateNewBook(ctx context.Context, book *pb.BookPayload) (*pb.Message, error) {
	if err := g.ServiceMongo.CreateBook(ctx, models.BookMongo{
		Pages:   int64(book.Pages),
		Year:    int64(book.Year),
		Title:   book.Title,
		Content: book.Content,
	}); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &pb.Message{
		Message: "Created",
		Error:   "",
	}, nil
}

func (g *grpcServer) UpdateByIdBook(ctx context.Context, updatePayload *pb.UpdateBook) (*pb.Message, error) {
	if err := g.ServiceMongo.UpdateBook(ctx, updatePayload.Id.ID, models.BookMongo{
		Pages:   int64(updatePayload.Book.Pages),
		Year:    int64(updatePayload.Book.Year),
		Title:   updatePayload.Book.Title,
		Content: updatePayload.Book.Content,
	}); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.Message{
		Message: "Updated",
		Error:   "",
	}, nil
}
