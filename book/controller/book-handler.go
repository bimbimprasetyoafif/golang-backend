package controller

import (
	"context"
	"net/http"
	"strconv"

	"example.com/book"

	"example.com/models"
	"github.com/labstack/echo"
)

var aBook models.Book

type BookHandler struct{
	BService book.Service
}

type ResponseError struct {
	Message string `json:"message"`
}

func NewBookHandler(e *echo.Echo, bs book.Service){
	handler := &BookHandler{
		BService: bs,
	}
	e.POST("/book/", handler.CreateBook)
	e.PUT("/book/:id", handler.UpdateBook)
	e.DELETE("/book/:id", handler.DeleteBook)
	e.GET("/book/:id", handler.GetById)
	e.GET("/book/", handler.GetAll)
}

func (b *BookHandler) CreateBook(ec echo.Context) error {

	err := ec.Bind(&aBook)
	if err != nil{
		return ec.JSON(http.StatusBadRequest,ResponseError{Message:"bad request"})
	}

	ctx := ec.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = b.BService.CreateBook(ctx, &aBook)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message:"internal server error"})
	}

	return ec.JSON(http.StatusCreated, aBook)
}

func (b *BookHandler) UpdateBook(ec echo.Context) error {
	IdParam, err := strconv.Atoi(ec.Param("id"))
	id := int64(IdParam)

	err = ec.Bind(&aBook)
	if err != nil{
		return ec.JSON(http.StatusBadRequest,ResponseError{Message:"bad request"})
	}

	ctx := ec.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = b.BService.UpdateBook(ctx,id,&aBook)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message:"internal server error"})
	}

	return ec.JSON(http.StatusCreated, aBook)

}

func (b *BookHandler) DeleteBook(ec echo.Context) error{
	IdParam, err := strconv.Atoi(ec.Param("id"))
	if err != nil {
		return ec.JSON(http.StatusBadRequest, ResponseError{Message:"bad parameter"})
	}

	ctx := ec.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id := int64(IdParam)
	err = b.BService.DeleteBook(ctx,id)
	if err != nil {
		return ec.JSON(http.StatusNotFound, ResponseError{Message:"data not found"})
	}

	return ec.NoContent(http.StatusNoContent)
}

func (b *BookHandler) GetById(ec echo.Context) error{

	IdParam, err := strconv.Atoi(ec.Param("id"))
	if err != nil {
		return ec.JSON(http.StatusBadRequest, ResponseError{Message:"bad parameter"})
	}

	ctx := ec.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id := int64(IdParam)
	err = b.BService.GetById(ctx, id, &aBook)
	if err != nil {
		return ec.JSON(http.StatusNotFound, ResponseError{Message:"not found"})
	}
	return  ec.JSON(http.StatusOK, aBook)
}

func (b *BookHandler) GetAll(ec echo.Context) error {
	ctx := ec.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	books, err := b.BService.GetAll(ctx)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message:"server error"})
	}
	return ec.JSON(http.StatusOK, books)
}