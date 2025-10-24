package metadata

import (
	"io"
	"strings"
)

type MetaDataHandler struct {
	root      *INode
	serilizer INodeTreeSerializer
}

func NewFileMetaDataHandler() *MetaDataHandler {
	return &MetaDataHandler{
		root: &INode{
			Name:  "",
			IsDir: true,
		},
		serilizer: NewGobSerializer(),
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

func (fileMetaDataHandler *MetaDataHandler) saveSnapshot(writer io.Writer) {
	fileMetaDataHandler.serilizer.Encode(fileMetaDataHandler.root, writer)
}
