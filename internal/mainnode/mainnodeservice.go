package mainnode

import (
	"context"
	"fmt"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/api"
)

type MainNodeService struct {
	pb.UnimplementedMainNodeServiceServer
}

// Register implements rpc.MetaDataServiceServer.
func (mds MainNodeService) Register(ctx context.Context, dataNodeInfo *pb.DataNodeInfo) (*pb.RegisterReply, error) {
	reply := pb.RegisterReply{Status: "SUCCESS"}
	fmt.Println(dataNodeInfo.NodeId)
	return &reply, nil
}
