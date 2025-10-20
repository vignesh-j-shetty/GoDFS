package datanode

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MainNodeConnectionManager interface {
	ConnectWithServer() error
}

type MainNodeConnectionManagerImpl struct {
	client pb.MainNodeServiceClient
}

func NewMainNodeConnector(config config.Config) (MainNodeConnectionManager, error) {
	fmt.Printf("Attempting to connect with %s\n", config.MainNodeUrl)
	conn, err := grpc.NewClient(config.MainNodeUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("CONNECTION TO METADATA SERVER FAILED WITH ERROR %w", err)
	}
	c := pb.NewMainNodeServiceClient(conn)

	return &MainNodeConnectionManagerImpl{
		client: c,
	}, nil
}

func (sc *MainNodeConnectionManagerImpl) ConnectWithServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 300)
	defer cancel()

	id := uuid.New()
	req := &pb.DataNodeInfo{NodeId: id.String(), RpcEndpoint: "0.0.0.0:8080"}
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
