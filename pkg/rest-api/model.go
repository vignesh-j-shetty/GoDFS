package restapi

type FolderRequest struct {
	Path string `json:"path"`
	FolderName string `json:"folderName"`
}

type FolderContentList struct {
	FolderContentList []string `json:"folderContentList"`
}

type FileCreateRequest struct {
	Path string `json:"path"`
	FileName string `json:"fileName"`
	Size uint64 `json:"fileSize"`
}

type Response struct {
	Status string `json:"status"`
	Error string `json:"error"`
	Data any `json:"data"`
}

type ChunkInfo struct {
	ChunkId string `json:"chunkId"`
	UploadUrl []string `json:"uploadUrl"`
}