package client

import "context"

type Client interface {
	Upload(ctx context.Context, localPath string, remotePath string) (fileID string, err error)
}