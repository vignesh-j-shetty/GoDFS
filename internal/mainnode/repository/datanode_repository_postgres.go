package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
)

type DataNodeRepositoryPostgres struct {
	conn   *pgx.Conn
	config config.Config
}

func NewDataNodeRepositoryPostgres(config config.Config) (*DataNodeRepositoryPostgres, error) {
	conn, err := pgx.Connect(context.Background(), config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return &DataNodeRepositoryPostgres{
		conn: conn,
		config: config,
	}, nil
}

func (repository DataNodeRepositoryPostgres) CreateDataNode(ctx context.Context, dataNode DataNode) error {
	txOptions := pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	}
	
	tx, err := repository.conn.BeginTx(ctx, txOptions)
	defer func() {
        if rbe := tx.Rollback(ctx); rbe != nil && rbe != pgx.ErrTxClosed {
            // Log the rollback error if it's not the expected "already closed" error
            log.Printf("Warning: Failed to rollback transaction: %v", rbe)
        }
    }()

	if err != nil {
		return fmt.Errorf("unable to create transaction %w", err)
	}

	var primaryCount int = 0
	if err := tx.QueryRow(ctx, "SELECT COUNT(*) FROM CHUNK_SERVER WHERE ROLE = 'PRIMARY'").Scan(&primaryCount); err != nil {
		return fmt.Errorf("%w", err)
	}

	var nodeType string
	if primaryCount < repository.config.DataNodeCount {
		nodeType = "PRIMARY"
	} else {
		nodeType = "REPLICA"
	}

	insertQuery := "INSERT INTO CHUNK_SERVER (SERVER_ID, RPC_ENDPOINT, ROLE) VALUES ($1, $2, $3)"
	_, err = tx.Exec(ctx, insertQuery, dataNode.NodeId, dataNode.RpcEndpoint, nodeType)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return ErrDuplicateDataNode
		}
		log.Printf("Unhandled db error %s ", err.Error())
		return fmt.Errorf("DATABASE ERROR : %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

	return nil
}

func (repository DataNodeRepositoryPostgres) UpdateRpcEndpoint(ctx context.Context, nodeID string, rcpEndpoint string) error {
	query := "UPDATE CHUNK_SERVER SET RPC_ENDPOINT = $1, WHERE SERVER_ID = $2"
	result, err := repository.conn.Exec(ctx, query, rcpEndpoint, nodeID)

	if err != nil {
		return fmt.Errorf("DATABASE ERROR : %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrDataNodeNotFound
	}

	return nil
}

func (repository DataNodeRepositoryPostgres) GetPrimaryDataNodeCount(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) from CHUNK_SERVER WHERE ROLE = 'PRIMARY'"
	var count int
	err := repository.conn.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		// Handle the case where an actual database error occurred (e.g., connection issue)
		return 0, fmt.Errorf("DATABASE ERROR getting primary data node count: %w", err)
	}
	return count, nil
}
