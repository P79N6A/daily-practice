package services

import (
	"fmt"
	"log"
	"net"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/services/gen"
	"google.golang.org/grpc"
)

func Serve(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("[GRPC] services.UserServer is listening on port: %d\n", port)

	grpcServer := grpc.NewServer()
	services.RegisterUserServiceServer(grpcServer, &UserServer{})
	grpcServer.Serve(lis)
}
