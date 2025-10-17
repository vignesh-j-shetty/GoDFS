package mainnode

import (
	"context"
	"fmt"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/api"
)

type MetaDataService struct {
	pb.UnimplementedMainNodeServiceServer
}

// Register implements rpc.MetaDataServiceServer.
func (mds MetaDataService) Register(ctx context.Context, chunkServerInfo *pb.DataNodeInfo) (*pb.RegisterReply, error) {
	reply := pb.RegisterReply{Status: "SUCCESS"}
	fmt.Println(chunkServerInfo.ServerId)
	return &reply, nil
}
