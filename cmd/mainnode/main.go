package main

import (
	"fmt"
	"log"
	"net"
	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load()
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Printf("%s", err.Error());
	}

	s := grpc.NewServer()
	mainNodeService, err := mainnode.NewMainNodeService()

	if err != nil {
		fmt.Printf("Error while starting MainNode Service %s", err.Error())
	}

	pb.RegisterMainNodeServiceServer(s, mainNodeService)
	fmt.Printf("Starting to listen at 8080")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}