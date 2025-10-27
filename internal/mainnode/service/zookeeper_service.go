package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
	"github.com/vignesh-j-shetty/GoDFS/pkg/zookeeper"
)

type ZookeeperService struct {
	conn *zk.Conn
	evCh <-chan zk.Event
	config config.Config
}

func NewZookeeperService(config config.Config) (*ZookeeperService, error) {
	conn, evCh, err := zk.Connect(config.ZookeepersServers, 10 * time.Second)

	if err != nil {
		return nil, err
	}

	return &ZookeeperService{
		conn: conn,
		evCh: evCh,
		config: config,
	}, nil
}

func (service *ZookeeperService) WatchLoop() error {
	// Ensure /chunkservers prefix path exits
	err := service.ensurePath(commonconstants.ChunkServerPrefixPath)
	if err != nil {
		return err
	}
	for {
		children, _, ch, err := service.conn.ChildrenW(commonconstants.ChunkServerPrefixPath)
		if err != nil {
			return err
		}
		var chunkServers []zookeeper.ChunkServerInfo
		for _, child := range children {
			fullPath := commonconstants.ChunkServerPrefixPath + "/" + child
			// Ignore stat
			data, _, err := service.conn.Get(fullPath)

			if err != nil {
				fmt.Printf("error while reading chunkserver details %s", err.Error())
				continue
			}

			var chunkServerInfo zookeeper.ChunkServerInfo
			json.Unmarshal(data, &chunkServerInfo)
			chunkServers = append(chunkServers, chunkServerInfo)
		}
		fmt.Printf("Size %d \n", len(chunkServers))
		for _, chunchunkServersInfo := range chunkServers {
			fmt.Printf("Upload Urls %s\n", chunchunkServersInfo.UploadUrl)
		}
		ev := <-ch
		fmt.Println(ev.State.String())
	}
}

func (s *ZookeeperService) ensurePath(path string) error {
	exists, _, err := s.conn.Exists(path)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	// Create persistent node
	_, err = s.conn.Create(path, []byte{}, 0, zk.WorldACL(zk.PermAll))
	if err == zk.ErrNodeExists {
		return nil
	}
	return err
}