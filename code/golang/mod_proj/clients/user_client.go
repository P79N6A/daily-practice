package main

import (
	"context"
	"log"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/services/gen"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":1551", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := services.NewUserServiceClient(conn)
	req := services.UserRequest{
		Name: "moye",
	}
	resp, err := client.GetUser(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
