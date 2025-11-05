package service

import (
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/metadata"
)

type MetaDataService struct {
	handler metadata.MetaDataHandler
	config config.Config
}

func NewMetaDataService(config config.Config) MetaDataService {
	return MetaDataService{
		handler: *metadata.NewFileMetaDataHandler(),
		config: config,
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

func (mds MetaDataService) CreateFiles(path string, fileName string, fileSize uint64) error {
	err := mds.handler.CreateFile(path, fileName, fileSize)
	if err != nil {
		return err
	}

	return nil
}