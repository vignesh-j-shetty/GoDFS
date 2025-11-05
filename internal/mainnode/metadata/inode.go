package metadata

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
	ChunkID uint64
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

func (inode *INode) CreateFile(fileName string, fileSize uint64) error {
	if !inode.IsDir {
		return ErrInvalidOperation
	}

	for _, child := range inode.Children {
		if child.Name == fileName {
			return ErrDuplicateName
		}
	}
	newChild := INode{
		Name:     fileName,
		IsDir:    false,
		Children: nil,
	}

	//id := uuid.New()

	inode.Children = append(inode.Children, newChild)
	return nil
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
