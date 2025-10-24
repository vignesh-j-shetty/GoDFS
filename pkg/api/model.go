package api

type FolderRequest struct {
	Path       string `json:"path"`
	FolderName string `json:"folderName"`
}

type FolderContentList struct {
	FolderContentList []string `json:"folderContentList"`
}