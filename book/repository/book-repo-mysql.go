package repository

import (
	"context"
	"database/sql"
	"log"

	"example.com/book"
	"example.com/models"
)

var tempBook models.Book

type MysqlRepositoryBook struct {
	Conn *sql.DB
}

func NewMysqlRepositoryBook(Conn *sql.DB) book.Repository {
	return &MysqlRepositoryBook{Conn}
}

func (db *MysqlRepositoryBook) CreateBook(ctx context.Context, m models.Book) error {

	query := "INSERT book SET pages = ?, year = ?, title = ?, content = ?"

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	res, err := stmt.ExecContext(ctx, m.Pages, m.Year, m.Title, m.Content)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	m.ID = id

	return nil
}

func (db *MysqlRepositoryBook) DeleteBook(ctx context.Context, id int64) error {
	query := "DELETE FROM book WHERE id = ?"
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = db.GetById(ctx, id, tempBook)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (db *MysqlRepositoryBook) UpdateBook(ctx context.Context, id int64, m models.Book) error {
	query := "UPDATE book SET pages = ? , year = ? , title = ? , content = ? WHERE id = ?"
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = db.GetById(ctx, id, tempBook)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, m.Pages, m.Year, m.Title, m.Content, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (db *MysqlRepositoryBook) GetById(ctx context.Context, id int64, res models.Book) error {
	query := "SELECT id, pages, year, title, content FROM book WHERE id = ?"
	row := db.Conn.QueryRowContext(ctx, query, id).Scan(&res.ID, &res.Pages, &res.Year, &res.Title, &res.Content)

	if row == sql.ErrNoRows {
		log.Println(row.Error())
		return row
	}

	return nil
}

func (db *MysqlRepositoryBook) GetAll(ctx context.Context) ([]models.Book, error) {
	query := "SELECT id, pages, year, title, content FROM book"
	rows, err := db.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	res := make([]models.Book, 0)
	var each models.Book
	for rows.Next() {
		err := rows.Scan(&each.ID, &each.Pages, &each.Year, &each.Title, &each.Content)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		res = append(res, each)
	}
	return res, nil
}
