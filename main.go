package main

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"

	"example.com/book/database"

	MyBookHandler "example.com/book/controller"
	MyBookRepo "example.com/book/repository"
	MyBookService "example.com/book/service"
)

func init(){

	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Cannot read :"+err.Error())
	}

}

func main()  {

	db := database.InitDb()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	br := MyBookRepo.NewMysqlRepositoryBook(db)
	bs := MyBookService.NewBookService(br, timeoutContext)

	MyBookHandler.NewBookHandler(e,bs)

	log.Fatal(e.Start(viper.GetString("server.host")))
	
}