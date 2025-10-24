package datanode

import (
	"context"
	"fmt"
	"time"
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
	config config.Config
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
		config: config,
	}, nil
}

func (sc *MainNodeConnectionManagerImpl) ConnectWithServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 300)
	defer cancel()

	
	req := &pb.DataNodeInfo{RpcEndpoint: sc.config.SelfUrl}
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
