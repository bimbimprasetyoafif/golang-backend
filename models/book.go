package models

import "gopkg.in/mgo.v2/bson"

type Book struct {
	ID      int64  `json:"id"`
	Pages   int64  `json:"pages"`
	Year    int64  `json:"year"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type BookMongo struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Pages   int64         `json:"pages" bson:"pages"`
	Year    int64         `json:"year" bson:"year"`
	Title   string        `json:"title" bson:"title"`
	Content string        `json:"content" bson:"content"`
}
