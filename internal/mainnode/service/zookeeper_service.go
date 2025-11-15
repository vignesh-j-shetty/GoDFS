package service

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/model"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
)

type ZookeeperService struct {
	conn   *zk.Conn
	evCh   <-chan zk.Event
	config config.Config
	DataNode[] model.DataNodeInfo
}

func NewZookeeperService(config config.Config) (*ZookeeperService, error) {
	conn, evCh, err := zk.Connect(config.ZookeepersServers, 10*time.Second)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connect with zookeeper successfully")
	return &ZookeeperService{
		conn:   conn,
		evCh:   evCh,
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
		var dataNodeServers []model.DataNodeInfo
		for _, child := range children {
			fullPath := commonconstants.ChunkServerPrefixPath + "/" + child
			// Ignore stat
			data, _, err := service.conn.Get(fullPath)

			if err != nil {
				fmt.Printf("error while reading chunkserver details %s", err.Error())
				continue
			}

			var chunkServerInfo model.DataNodeInfo
			json.Unmarshal(data, &chunkServerInfo)
			dataNodeServers = append(dataNodeServers, chunkServerInfo)
		}

		sort.Slice(dataNodeServers, func(i, j int) bool {
			return dataNodeServers[i].FreeSpace > dataNodeServers[j].FreeSpace
		})
		service.DataNode = dataNodeServers
		<-ch
		fmt.Println("Datanode info updated")
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

func (s *ZookeeperService) GetActiveDatanodes() []model.DataNodeInfo {
	return s.DataNode
}