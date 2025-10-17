package mainnode

import (
	"context"
	"os"

	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/repository"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/api"
)

type MainNodeServiceImpl struct {
	repository repository.DataNodeRepository
	pb.UnimplementedMainNodeServiceServer
}

func NewMainNodeService() (*MainNodeServiceImpl, error) {
	repo, err := repository.NewDataNodeRepositoryPostgres(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &MainNodeServiceImpl{
		repository: repo,
	}, nil
}

// Register implements rpc.MetaDataServiceServer.
func (mds MainNodeServiceImpl) Register(ctx context.Context, dataNodeInfo *pb.DataNodeInfo) (*pb.RegisterReply, error) {
	data := repository.DataNode {
		NodeId: dataNodeInfo.NodeId,
		RpcEndpoint: dataNodeInfo.RpcEndpoint,
		Role: "PRIMARY",
	}
	err := mds.repository.InsertDataNode(ctx, data)
	if err != nil {
		return nil, err
	}
	reply := pb.RegisterReply{Status: "SUCCESS"}
	return &reply, nil
}
