package main

import (
	"fmt"
	"log"
	"net"
	"github.com/vignesh-j-shetty/GoDFS/internal/metadataserver"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/rpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5151")
	if err != nil {
		fmt.Printf("%s", err.Error());
	}
	s := grpc.NewServer()
	pb.RegisterMetaDataServiceServer(s, metadataserver.MetaDataService {})
	fmt.Printf("Starting to listen at 5151")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}