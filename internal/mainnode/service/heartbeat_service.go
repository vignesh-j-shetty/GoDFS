package service

import (
	"context"
	"fmt"

	pb "github.com/vignesh-j-shetty/GoDFS/pkg/rpc"
)

type HeartBeatService struct {
	pb.UnimplementedHeartBeatServiceServer
}

func (service *HeartBeatService) Heartbeat(ctx context.Context, in *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	fmt.Println("Recevied Heart beat")

	return &pb.HeartbeatResponse{
		Status: "Sucess",
	}, nil
}