package service

import (
	"fmt"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
)

const zkPathPrefix = "/chunkservers"

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
	err := service.ensurePath(zkPathPrefix)
	if err != nil {
		return err
	}
	for {
		children, _, ch, err := service.conn.ChildrenW(zkPathPrefix)
		if err != nil {
			return err
		}

		fmt.Println("Children ", children)

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