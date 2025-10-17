package datanode

import (
	"context"
	"fmt"
	"time"

	pb "github.com/vignesh-j-shetty/GoDFS/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/google/uuid"
)

type MainNodeConnectionManager interface {
	ConnectWithServer() error
}

type MainNodeConnectionManagerImpl struct {
	client pb.MainNodeServiceClient
}

func NewServerConnector() (MainNodeConnectionManager, error) {
	conn, err := grpc.NewClient(":5151", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("CONNECTION TO METADATA SERVER FAILED WITH ERROR %w", err)
	}
	c := pb.NewMainNodeServiceClient(conn)

	return &MainNodeConnectionManagerImpl {
		client: c,
	}, nil
}

func (sc *MainNodeConnectionManagerImpl) ConnectWithServer() error {
	id := uuid.New()
	req := &pb.DataNodeInfo {ServerId: id.String(), RpcEndpoint: "0.0.0.0:8080"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := sc.client.Register(ctx, req)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("RETURN RPC FAILED: Context deadline exceeded %w", err)
        }
		return fmt.Errorf("RETURN RPC FAILED %w", err)
	}
	
	println(reply.GetStatus())
	return nil
}