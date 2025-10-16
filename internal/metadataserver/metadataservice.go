package metadataserver

import (
	"context"
	"fmt"

	pb "github.com/vignesh-j-shetty/GoDFS/pkg/rpc"
)

type MetaDataService struct {
	pb.UnimplementedMetaDataServiceServer
}

func (s *MetaDataService) Connect(cxt context.Context, connect_request *pb.ConnectRequest) {
	id := connect_request.GetServerId()
	fmt.Println(id)
}