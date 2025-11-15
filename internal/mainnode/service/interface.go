package service

import "github.com/vignesh-j-shetty/GoDFS/internal/mainnode/model"

type ActiveDatanodeProvider interface {
	// GetActiveDatanodes returns a slice of active DataNodeInfo
	GetActiveDatanodes() []model.DataNodeInfo
}
