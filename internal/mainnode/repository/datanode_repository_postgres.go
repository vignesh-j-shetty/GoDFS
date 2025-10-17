package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type DataNodeRepositoryPostgres struct {
	conn *pgx.Conn
}

func NewDataNodeRepositoryPostgres(databaseURL string) (*DataNodeRepositoryPostgres, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	return &DataNodeRepositoryPostgres{
		conn: conn,
	}, nil
}

func (repository DataNodeRepositoryPostgres) InsertDataNode(ctx context.Context, dataNode DataNode) error {
	query := "INSERT INTO CHUNK_SERVER (SERVER_ID, RPC_ENDPOINT, ROLE) VALUES ($1, $2, $3)"
	_, err := repository.conn.Exec(ctx, query, dataNode.NodeId, dataNode.RpcEndpoint, dataNode.Role)
	return err
}