package service

import (
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
	datanodeallocator "github.com/vignesh-j-shetty/GoDFS/internal/mainnode/datanode-allocator"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/metadata"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/model"
)

type MetaDataService struct {
	handler metadata.MetaDataHandler
	config config.Config
	activeDatanodeProvider ActiveDatanodeProvider
}

func NewMetaDataService(config config.Config, activeDatanodeProvider ActiveDatanodeProvider) MetaDataService {
	return MetaDataService{
		handler: *metadata.NewFileMetaDataHandler(config),
		config: config,
		activeDatanodeProvider: activeDatanodeProvider,
	}
}

func (mds *MetaDataService) CreateFolder(path string, folderName string) error {
	err := mds.handler.CreateFolder(path, folderName)

	if err != nil {
		return err
	}

	return nil
}


func (mds *MetaDataService) Delete(path string) error {
	err := mds.handler.Delete(path)
	if err != nil {
		return err
	}

	return nil
}

func (mds MetaDataService) GetFolderContents(path string) ([]string, error) {
	return mds.handler.GetFolderContents(path)
}

func (mds MetaDataService) CreateFiles(path string, fileName string, fileSize uint64) ([]model.ChunkLocationInfo, error) {
	fileMetaData, err := mds.handler.CreateFile(path, fileName, fileSize)
	if err != nil {
		return nil, err
	}
	var chunkInfo [] model.ChunkInfo
	for _, chunk := range fileMetaData {
		chunkInfo = append(chunkInfo, model.ChunkInfo{
			ID:   chunk.ChunkID,
			Size: chunk.Size,
		})
	}

	activeNodes := mds.activeDatanodeProvider.GetActiveDatanodes()

	chunkAllocationInfo, err := datanodeallocator.AllocateDataNode(activeNodes, chunkInfo, 3)
	if err != nil {
		return nil, err
	}
	return chunkAllocationInfo, nil
}