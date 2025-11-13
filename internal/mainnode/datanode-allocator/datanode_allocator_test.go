package datanodeallocator

import (
	"fmt"
	"testing"

	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/model"
	"github.com/vignesh-j-shetty/GoDFS/pkg/datanode"
)

func TestAllocateDataNode_Success(t *testing.T) {
	datanodes := []datanode.DataNodeInfo{
		{Id: "dn1", FreeSpace: 100},
		{Id: "dn2", FreeSpace: 200},
		{Id: "dn3", FreeSpace: 300},
	}
	chunks := []model.ChunkInfo{
		{ID: "chunk1", Size: 50},
		{ID: "chunk2", Size: 100},
	}
	replicationFactor := uint16(2)
	locs, err := AllocateDataNode(datanodes, chunks, replicationFactor)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for _, loc := range locs {
		fmt.Println("id " + loc.ChunkID, " datanodes ", loc.DataNodeIDs)
	}

	if len(locs) != len(chunks) {
		t.Errorf("expected %d chunk locations, got %d", len(chunks), len(locs))
	}
	for _, loc := range locs {
		if len(loc.DataNodeIDs) != int(replicationFactor) {
			t.Errorf("expected %d replicas, got %d", replicationFactor, len(loc.DataNodeIDs))
		}
	}
}

func TestAllocateDataNode_NotEnoughSpace(t *testing.T) {
	datanodes := []datanode.DataNodeInfo{
		{Id: "dn1", FreeSpace: 10},
		{Id: "dn2", FreeSpace: 10},
	}
	chunks := []model.ChunkInfo{
		{ID: "chunk1", Size: 50},
	}
	replicationFactor := uint16(2)
	_, err := AllocateDataNode(datanodes, chunks, replicationFactor)
	if err == nil {
		t.Fatalf("expected error due to insufficient space, got nil")
	}
}
