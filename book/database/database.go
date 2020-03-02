package database

import (

	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"

)

func InitDb () (db *sql.DB){
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		viper.GetString("database.mysql.username"),
		viper.GetString("database.mysql.password"),
		viper.GetString("database.mysql.hostname"),
		viper.GetString("database.mysql.port"),
		viper.GetString("database.mysql.database"),
	)
	db, err := sql.Open(viper.GetString("database.mysql.driver"),connect)
	if err != nil {
		log.Println("Cannot connect database :"+err.Error())
	}
	return db
}