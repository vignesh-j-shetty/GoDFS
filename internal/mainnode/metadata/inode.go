package metadata

import "github.com/google/uuid"

// import (
// 	"github.com/google/uuid"
// )

// This represents single File or Directory
type INode struct {
	Name  string
	IsDir bool
	// len(Children) == 0 when IsDir is False
	Children []INode
	// len(FileMetaData) == 0 when IsDir is True
	FileMetaData []FileMetaData
}

type FileMetaData struct {
	ChunkID string
	Size    uint32
}

func (inode *INode) CreateFolder(folderName string) error {
	if !inode.IsDir {
		return ErrInvalidOperation
	}

	for _, child := range inode.Children {
		if child.Name == folderName {
			return ErrDuplicateName
		}
	}

	inode.Children = append(inode.Children, INode{
		Name:     folderName,
		IsDir:    true,
		Children: nil,
	})
	return nil
}

func (inode *INode) CreateFile(fileName string, fileSize uint64, MaxChunkSize uint64) (INode, error) {
	if !inode.IsDir {
		return INode{}, ErrInvalidOperation
	}

	for _, child := range inode.Children {
		if child.Name == fileName {
			return INode{}, ErrDuplicateName
		}
	}
	newChild := INode{
		Name:     fileName,
		IsDir:    false,
		Children: nil,
	}

	chunkCount := fileSize / MaxChunkSize

	for range chunkCount {
		id := uuid.New().String()
		newChild.FileMetaData = append(newChild.FileMetaData, FileMetaData{
			ChunkID: id,
			Size:    uint32(MaxChunkSize),
		})
	}

	remainingSize := fileSize % MaxChunkSize
	if remainingSize > 0 {
		id := uuid.New().String()
		newChild.FileMetaData = append(newChild.FileMetaData, FileMetaData{
			ChunkID: id,
			Size:    uint32(remainingSize),
		})
	}
	
	inode.Children = append(inode.Children, newChild)
	return newChild, nil
}

func (inode *INode) DeleteChild(fileName string) error {

	if !inode.IsDir {
		return ErrInvalidOperation
	}

	targetIndex := -1

	for i, child := range inode.Children {
		if child.Name == fileName {
			if child.IsDir && len(child.Children) > 0 {
				return ErrFolderNotEmpty
			}
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		return ErrFileNotFound
	}

	inode.Children = append(inode.Children[:targetIndex], inode.Children[targetIndex+1:]...)
	return nil
}
