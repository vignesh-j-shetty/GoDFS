package datanode

type ChunkServerInfo struct {
	Id string `json:"id"`
	UploadUrl string `json:"uploadUrl"`
	FreeSpace uint64 `json:"freeSpace"`
}