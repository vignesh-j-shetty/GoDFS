package datanodeallocator

import (
	"sort"

	"errors"

	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/model"
	"github.com/vignesh-j-shetty/GoDFS/pkg/datanode"
)

func AllocateDataNode(datanodeinfos []datanode.DataNodeInfo, chunks []model.ChunkInfo, replicationFactor uint16) ([]model.ChunkLocationInfo, error) {
	var chunkLocations []model.ChunkLocationInfo
	numDataNodes := len(datanodeinfos)
	// Sort datanodes by free space descending
	sort.SliceStable(datanodeinfos, func(i, j int) bool {
		return datanodeinfos[i].FreeSpace > datanodeinfos[j].FreeSpace
	})
	for _, chunk := range chunks {
		var assignedIDs []string
		count := 0
		for i := 0; i < numDataNodes && count < int(replicationFactor); i++ {
			if datanodeinfos[i].FreeSpace >= uint64(chunk.Size) {
				assignedIDs = append(assignedIDs, datanodeinfos[i].Id)
				count++
			}
		}
		if len(assignedIDs) < int(replicationFactor) {
			return nil, errors.New("not enough free space to allocate chunk " + chunk.ID)
		}
		chunkLoc := model.ChunkLocationInfo{
			ChunkID:     chunk.ID,
			DataNodeIDs: assignedIDs,
		}
		chunkLocations = append(chunkLocations, chunkLoc)
	}
	return chunkLocations, nil
}
