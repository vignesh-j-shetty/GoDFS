package main

import (
	"log"
	"net"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/handler"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/service"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/rpc"
	"google.golang.org/grpc"
)

func main() {
	// Load env variables
	_ = godotenv.Load()

	// Gin setup
	r := gin.Default()
	folderOps := handler.NewFolderOpsHandler()
	folderOps.InitRoutes(r)

	// Grpc setup 
	grpcAddr := ":50051"
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", grpcAddr, err)
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterHeartBeatServiceServer(grpcSrv, &service.HeartBeatService {})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		log.Printf("https server listening on 8080")
		r.Run(":8080")
	}()

	go func() {
		defer wg.Done()
		log.Printf("gRPC server listening on %s", grpcAddr)
		if err := grpcSrv.Serve(lis); err != nil {
			log.Printf("Error in starting gRPC server %s\n", err.Error())
		}
	}()
	wg.Wait()
}