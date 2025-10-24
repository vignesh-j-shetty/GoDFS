package service

import "github.com/vignesh-j-shetty/GoDFS/internal/mainnode/metadata"

type MetaDataService struct {
	handler metadata.MetaDataHandler
}

func NewMetaDataService() MetaDataService {
	return MetaDataService{
		handler: *metadata.NewFileMetaDataHandler(),
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