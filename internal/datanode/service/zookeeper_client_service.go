package service

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/go-zookeeper/zk"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
	"github.com/vignesh-j-shetty/GoDFS/pkg/datanode"
	"github.com/vignesh-j-shetty/GoDFS/pkg/platform"
)

type ZookeeperClientService struct {
	conn *zk.Conn
	evCh <-chan zk.Event
	config config.Config
}


func NewZookeeperClientService(config config.Config) (*ZookeeperClientService, error) {
	conn, evCh, err := zk.Connect(config.ZookeepersServers, time.Second)

	if err != nil {
		return nil, err
	}

	return &ZookeeperClientService{
		conn: conn,
		evCh: evCh,
		config: config,
	}, nil
}

func (s *ZookeeperClientService) Register() error {
	freeSpace, err := platform.GetFreeSpace(s.config.ChunkFilePath)

	if err != nil {
		return err
	}

	chunkInfo := datanode.ChunkServerInfo {
		Id: s.config.Id,
		UploadUrl: s.config.UploadUrl,
		FreeSpace: freeSpace,
	}
	jsonData, err := json.Marshal(chunkInfo)

	if err != nil {
		return err
	}

	created, err := s.conn.Create(commonconstants.ChunkServerPrefixPath + "/" + chunkInfo.Id, jsonData, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}
	fmt.Printf("created %s", created)
	return nil
}