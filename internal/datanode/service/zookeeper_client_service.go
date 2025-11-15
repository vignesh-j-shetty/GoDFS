package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/storage"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/model"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
)

type ZookeeperClientService struct {
	conn           *zk.Conn
	evCh           <-chan zk.Event
	config         config.Config
	storageHandler storage.StorageHandler
}

func NewZookeeperClientService(config config.Config) (*ZookeeperClientService, error) {
	conn, evCh, err := zk.Connect(config.ZookeepersServers, time.Second)

	if err != nil {
		return nil, err
	}

	return &ZookeeperClientService{
		conn:           conn,
		evCh:           evCh,
		config:         config,
		storageHandler: storage.NewStorageHandler(config),
	}, nil
}

func (s *ZookeeperClientService) Register() error {
	freeSpace, err := s.storageHandler.GetFreeSpace()

	if err != nil {
		fmt.Println("Failed to get disk info ", err.Error(), " for path ", s.config.ChunkFilePath)
		return err
	}

	fmt.Println("Disk size :", freeSpace)

	chunkInfo := model.DataNodeInfo{
		Id:        s.config.Id,
		UploadUrl: s.config.UploadUrl,
		FreeSpace: freeSpace,
	}
	jsonData, err := json.Marshal(chunkInfo)

	if err != nil {
		return err
	}

	created, err := s.conn.Create(commonconstants.ChunkServerPrefixPath+"/"+chunkInfo.Id, jsonData, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}
	fmt.Printf("created %s", created)
	return nil
}
