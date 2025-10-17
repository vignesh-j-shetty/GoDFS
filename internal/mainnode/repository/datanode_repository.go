package repository

import "context"

type DataNode struct {
	NodeId      string
	RpcEndpoint string
	Role        string
}

type DataNodeRepository interface {
	InsertDataNode(context.Context, DataNode) error
}