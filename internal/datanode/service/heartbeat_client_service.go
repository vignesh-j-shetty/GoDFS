package service

import (
	"fmt"
	"time"

	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	pb "github.com/vignesh-j-shetty/GoDFS/pkg/rpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HeartbeatClientService struct {
	done chan struct{}
}

func (service *HeartbeatClientService) StartHeartbeatService(config config.Config, ctx context.Context) error {
	service.done = make(chan struct{})
	conn, err := grpc.NewClient(config.MainNodeRCPUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	client := pb.NewHeartBeatServiceClient(conn)

	ticker := time.NewTicker(time.Second)
	
	// Create Thread
	go func() {
		defer func ()  {
			ticker.Stop()
			conn.Close()
			close(service.done)
		}()
		for {
			select {
				case <- ticker.C:
					service.sendHeartbeat(client)
				case <- ctx.Done():
					return
				}
			}
		}()
	return nil
}

func (service *HeartbeatClientService) Wait() {
	if service.done != nil {
		<-service.done
	}
}

func (service *HeartbeatClientService) sendHeartbeat(client pb.HeartBeatServiceClient) {
	req := &pb.HeartbeatRequest{
		ChunkIDs: []string{"chunk1", "chunk2", "chunk-xyz"},
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	resp, _ := client.Heartbeat(ctx, req)
	fmt.Println(resp.Status)
}