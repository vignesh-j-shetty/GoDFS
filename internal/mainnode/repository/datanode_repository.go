package repository

import (
	"context"
	"errors"
)

var (
    ErrDuplicateDataNode      = errors.New("data node already exists")
    ErrDataNodeNotFound       = errors.New("data node not found") 
)

type DataNode struct {
	NodeId      string
	RpcEndpoint string
	Role        string
}

type DataNodeRepository interface {
	// Inserts/Creates new entry in DataNode table
	CreateDataNode(context.Context, DataNode) error
	// Updates DataNode row
	UpdateRpcEndpoint(ctx context.Context, nodeID string, rcpEndpoint string) error
}