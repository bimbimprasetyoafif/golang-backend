package repository

import (
	"context"
	"database/sql"
	//"log"
	"testing"

	"example.com/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

var aBook models.Book

func initMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestGetBookId(t *testing.T) {

	db, mock := initMock(t)
	rows := sqlmock.NewRows([]string{
		"id",
		"pages",
		"year",
		"title",
		"content"}).AddRow(1, 123, 2019, "Title 1", "Just Content")

	query := "SELECT id, pages, year, title, content FROM book WHERE id = ?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	tDb := NewMysqlRepositoryBook(db)

	tId := int64(1)

	err := tDb.GetById(context.Background(), tId, aBook)

	assert.NoError(t, err)
	assert.NotNil(t, aBook)
}

func TestDeleteBook(t *testing.T) {
	db, mock := initMock(t)
	query := "DELETE FROM book WHERE id = ?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 0))

	tDb := NewMysqlRepositoryBook(db)
	err := tDb.DeleteBook(context.Background(), int64(1))

	assert.NoError(t, err)
}

//func TestUpdateBook (t *testing.T){
//	aBook := models.Book{
//		ID:      1,
//		Pages:   2,
//		Year:    2109,
//		Title:   "T",
//		Content: "C",
//	}
//	db, mock := initMock(t)
//	query := "UPDATE book SET pages = ? , year = ? , title = ? , content = ? WHERE id = ?"
//	prep := mock.ExpectPrepare(query)
//	prep.ExpectExec().WithArgs(aBook.Pages,aBook.Year, aBook.Title, aBook.Content, aBook.ID).WillReturnResult(sqlmock.NewResult(1,1))
//
//	tDb := NewMysqlRepositoryBook(db)
//	err := tDb.UpdateBook(context.Background(),1,aBook)
//
//	log.Print(err.Error())
//	t.Errorf(err.Error())
//	assert.NoError(t, err)
//}
