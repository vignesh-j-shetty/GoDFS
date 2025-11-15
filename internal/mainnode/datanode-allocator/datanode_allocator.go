package datanodeallocator

import (
	"errors"
	"fmt"
	"sort"

	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/model"
)

var ErrNotEnoughDatanodes = errors.New("not enough datanodes available for replication")
var ErrNotEnoughStorage = errors.New("not enough storage available for replication")

func AllocateDataNode(datanodeinfos []model.DataNodeInfo, chunks []model.ChunkInfo, replicationFactor uint16) ([]model.ChunkLocationInfo, error) {
	var chunkLocations []model.ChunkLocationInfo

	if len(datanodeinfos) < int(replicationFactor) {
		fmt.Println("Available Datnodes ", len(datanodeinfos), " Replication Factor ", replicationFactor)
		return nil, ErrNotEnoughDatanodes
	}

	for _, chunk := range chunks {

		var datanodeNodes []model.DataNodeInfo
		for i := range replicationFactor {
			// replicationFactor is guaranteed to be less than len(datanodeinfos)
			selectedNode := datanodeinfos[i]
			if selectedNode.FreeSpace < uint64(chunk.Size) {
				return nil, ErrNotEnoughStorage
			}
			datanodeNodes = append(datanodeNodes, selectedNode)
			datanodeinfos[i].FreeSpace -= uint64(chunk.Size)
		}
		
		sort.Slice(datanodeinfos, func(i, j int) bool {
			return datanodeinfos[i].FreeSpace > datanodeinfos[j].FreeSpace
		})

		chunkLocations = append(chunkLocations, model.ChunkLocationInfo{
			ChunkID:     chunk.ID,
			DataNodeIDs: datanodeNodes,
		})
	}
	return chunkLocations, nil
}
