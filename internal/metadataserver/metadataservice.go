package metadataserver

import (
	"context"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/rpc"
)

type MetaDataService struct {
	pb.UnimplementedMetaDataServiceServer
}

// Register implements rpc.MetaDataServiceServer.
func (mds MetaDataService) Register(context.Context, *pb.ChunkServerInfo) (*pb.RegisterReply, error) {
	reply := pb.RegisterReply{Status: "SUCCESS"}
	return &reply, nil
}
