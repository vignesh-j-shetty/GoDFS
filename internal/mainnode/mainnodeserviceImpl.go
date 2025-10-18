package mainnode

import (
	"context"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/repository"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/api"
)

type MainNodeServiceImpl struct {
	repository repository.DataNodeRepository
	pb.UnimplementedMainNodeServiceServer
}

func NewMainNodeService() (*MainNodeServiceImpl, error) {
	config := config.InitConfig()
	repo, err := repository.NewDataNodeRepositoryPostgres(config)
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
		Role: "",
	}

	err := mds.repository.CreateDataNode(ctx, data)
	if err != nil {
		if err == repository.ErrDuplicateDataNode {
			err = mds.repository.UpdateRpcEndpoint(ctx, data.NodeId, data.RpcEndpoint)
			if err != nil {
				return &pb.RegisterReply{Status: "FAILURE", ErrorMsg: err.Error()}, nil
			}
		}
		return nil, err
	}
	return &pb.RegisterReply{Status: "SUCCESS"}, nil
}