package mainnode

import (
	"context"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/api"
)

type MainNodeServiceImpl struct {
	pb.UnimplementedMainNodeServiceServer
}

func NewMainNodeService() (*MainNodeServiceImpl, error) {
	return &MainNodeServiceImpl{}, nil
}

// Register implements rpc.MetaDataServiceServer.
func (mds MainNodeServiceImpl) Register(ctx context.Context, dataNodeInfo *pb.DataNodeInfo) (*pb.RegisterReply, error) {

	return nil, nil
}