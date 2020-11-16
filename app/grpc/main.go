package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"

	"example.com/book/database"

	MyBookHandler "example.com/book/controller"
	MyBookRepo "example.com/book/repository"
	MyBookService "example.com/book/service"
	pb "example.com/pb"
)

func init() {

	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Cannot read :" + err.Error())
	}

}

func main() {
	db := database.InitDbMongo()
	port := viper.GetString("server.rpc")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	br := MyBookRepo.NewMongoRepositoryBook(db)
	bs := MyBookService.NewBookMongoService(br, timeoutContext)
	handler := MyBookHandler.NewGrpc(bs)
	s := grpc.NewServer()
	pb.RegisterBookGrpcServer(s, handler)

	run := fmt.Sprintf("GRPC Running On %s", port)
	fmt.Println(run)
	if err := s.Serve(lis); err != nil {
		fmt.Println("cannot run")
	}
}
