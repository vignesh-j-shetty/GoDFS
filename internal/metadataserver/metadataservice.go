package metadataserver

import (
	"context"
	"fmt"

	pb "github.com/vignesh-j-shetty/GoDFS/pkg/rpc"
)

type MetaDataService struct {
	pb.UnimplementedMetaDataServiceServer
}

// Register implements rpc.MetaDataServiceServer.
func (mds MetaDataService) Register(ctx context.Context, chunkServerInfo *pb.ChunkServerInfo) (*pb.RegisterReply, error) {
	reply := pb.RegisterReply{Status: "SUCCESS"}
	fmt.Println(chunkServerInfo.ServerId)
	return &reply, nil
}
