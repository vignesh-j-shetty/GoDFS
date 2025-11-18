package metadata

import (
	"io"
	"strings"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
)

type MetaDataHandler struct {
	root      *INode
	serilizer INodeTreeSerializer
	config	config.Config
}

func NewFileMetaDataHandler(config config.Config) *MetaDataHandler {
	return &MetaDataHandler{
		root: &INode{
			Name:  "",
			IsDir: true,
		},
		serilizer: NewGobSerializer(),
		config:    config,
	}
}

func NewFileMetaDataFromSnapshot(reader io.Reader) *MetaDataHandler {
	serializer := NewGobSerializer()
	var inode INode
	serializer.Decode(&inode, reader)
	return &MetaDataHandler{
		root:      &inode,
		serilizer: serializer,
	}
}

func (fileMetaDataHandler *MetaDataHandler) GetINodeFromPath(path string) (*INode, error) {
	if path == "/" || path == "" {
		return fileMetaDataHandler.root, nil
	}

	components := strings.Split(strings.Trim(path, "/"), "/")

	current := fileMetaDataHandler.root
	for _, name := range components {
		if !current.IsDir {
			return nil, ErrNotFolder
		}

		found := false
		for i := range current.Children {
			child := &current.Children[i]
			if child.Name == name {
				current = child
				found = true
				break
			}
		}

		if !found {
			return nil, ErrFileNotFound
		}
	}
	return current, nil
}

func (fileMetaDataHandler *MetaDataHandler) CreateFolder(path string, folderName string) error {
	parentINode, err := fileMetaDataHandler.GetINodeFromPath(path)

	if err != nil {
		return err
	}

	return parentINode.CreateFolder(folderName)
}

func (fileMetaDataHandler *MetaDataHandler) CreateFile(path string, fileName string, fileSize uint64) ([]FileMetaData, error) {
	parentINode, err := fileMetaDataHandler.GetINodeFromPath(path)

	if err != nil {
		return nil, err
	}

	fileNode, err := parentINode.CreateFile(fileName, fileSize, fileMetaDataHandler.config.ChunkSize)

	if err != nil {
		return nil, err
	}
	
	return fileNode.FileMetaData, nil
}

func (fileMetaDataHandler *MetaDataHandler) Delete(path string) error {
	lastIndex := strings.LastIndex(path, "/")

	if lastIndex == -1 {
		return ErrInvalidPath
	}

	parentPath := path[:lastIndex]
	nameToDelete := path[lastIndex + 1:]
	parentINode, err := fileMetaDataHandler.GetINodeFromPath(parentPath)

	if err != nil {
		return err
	}

	return parentINode.DeleteChild(nameToDelete)
}

func (fileMetaDataHandler *MetaDataHandler) GetFolderContents(path string) ([]string, error) {
	children, err := fileMetaDataHandler.GetINodeFromPath(path)

	if err != nil {
		return nil, err
	}

	size := len(children.Children)
	content := make([]string, size)
	for i, child := range children.Children {
		content[i] = child.Name
	}
	return content, nil
}

func (fileMetaDataHandler *MetaDataHandler) GetFileMetaData(path string) ([]FileMetaData, error) {
	inode, err := fileMetaDataHandler.GetINodeFromPath(path)
	if err != nil {
		return nil, err
	}
	if inode.IsDir {
		return nil, ErrInvalidOperation
	}
	return inode.FileMetaData, nil
}
