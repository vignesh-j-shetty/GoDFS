package chunkserver

import (
	"context"
	"fmt"
	"time"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/rpc"
	"google.golang.org/grpc"
)

type ServerConnectionManager interface {
	ConnectWithServer() error
}

type ServerConnectionManagerImpl struct {
	ctx context.Context
	client pb.MetaDataServiceClient
	cancel context.CancelFunc
}

func NewServerConnector() (ServerConnectionManager, error) {
	conn, err := grpc.NewClient(":5151")
	if err != nil {
		return nil, fmt.Errorf("CONNECTION TO METADATA SERVER FAILED WITH ERROR %w", err)
	}
	c := pb.NewMetaDataServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return &ServerConnectionManagerImpl {
		ctx: ctx,
		client: c,
		cancel: cancel,
	}, nil
}

func (sc *ServerConnectionManagerImpl) ConnectWithServer() error {
	req := &pb.ChunkServerInfo {ServerId: "jcdshhvb"}
	reply, err := sc.client.Register(sc.ctx, req)

	if err != nil {
		return fmt.Errorf("RETURN RPC FAILED %w", err)
	}

	println("%s", reply.GetStatus())
	return nil
}