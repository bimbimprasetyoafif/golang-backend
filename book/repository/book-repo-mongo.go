package repository

import (
	"context"
	"errors"
	"fmt"

	"example.com/book"
	"example.com/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoRepositoryBook struct {
	Conn *mgo.Session
}

const (
	DB_NAME       = "test"
	DB_COLLECTION = "book"
)

func NewMongoRepositoryBook(Conn *mgo.Session) book.RepositoryMongo {
	return &MongoRepositoryBook{Conn}
}

func (db *MongoRepositoryBook) CreateBook(ctx context.Context, m models.BookMongo) error {

	err := db.Conn.DB(DB_NAME).C(DB_COLLECTION).Insert(&m)
	if err != nil {
		return err
	}

	return nil
}

func (db *MongoRepositoryBook) DeleteBook(ctx context.Context, id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("id error")
	}

	oid := bson.ObjectIdHex(id)
	err := db.Conn.DB(DB_NAME).C(DB_COLLECTION).RemoveId(oid)
	if err != nil {
		return err
	}

	return nil
}

func (db *MongoRepositoryBook) UpdateBook(ctx context.Context, id string, m models.BookMongo) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("id error")
	}

	oid := bson.ObjectIdHex(id)
	err := db.Conn.DB(DB_NAME).C(DB_COLLECTION).UpdateId(oid, &m)
	if err != nil {
		return err
	}

	return nil
}

func (db *MongoRepositoryBook) GetById(ctx context.Context, id string, res models.BookMongo) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("id error")
	}

	oid := bson.ObjectIdHex(id)
	err := db.Conn.DB(DB_NAME).C(DB_COLLECTION).FindId(oid).One(&res)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoRepositoryBook) GetAll(ctx context.Context) ([]models.BookMongo, error) {
	var res []models.BookMongo
	err := db.Conn.DB(DB_NAME).C(DB_COLLECTION).Find(nil).All(&res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return res, nil
}
