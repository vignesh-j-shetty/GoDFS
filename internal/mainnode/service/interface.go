package service

import "github.com/vignesh-j-shetty/GoDFS/pkg/datanode"

type ActiveDatanodeProvider interface {
	// GetActiveDatanodes returns a slice of active DataNodeInfo
	GetActiveDatanodes() []datanode.DataNodeInfo
}
