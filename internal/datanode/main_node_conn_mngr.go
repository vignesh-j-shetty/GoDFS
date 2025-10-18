package datanode

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
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

func NewMainNodeConnector() (MainNodeConnectionManager, error) {
	conn, err := grpc.NewClient("mainnode:5151", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("CONNECTION TO METADATA SERVER FAILED WITH ERROR %w", err)
	}
	c := pb.NewMainNodeServiceClient(conn)

	return &MainNodeConnectionManagerImpl{
		client: c,
	}, nil
}

func (sc *MainNodeConnectionManagerImpl) ConnectWithServer() error {
	id := uuid.New()
	req := &pb.DataNodeInfo{NodeId: id.String(), RpcEndpoint: "0.0.0.0:8080"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 300)
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
