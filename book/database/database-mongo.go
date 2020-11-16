package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

func InitDbMongo() *mgo.Session {
	s, err := mgo.Dial(fmt.Sprintf("%s://%s",
		viper.GetString("database.mongo.driver"),
		viper.GetString("database.mongo.hostname")))

	if err != nil {
		panic(err)
	}

	return s
}
