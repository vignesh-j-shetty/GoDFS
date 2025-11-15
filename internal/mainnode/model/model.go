package model

type DataNodeInfo struct {
	Id        string `json:"id"`
	UploadUrl string `json:"uploadUrl"`
	FreeSpace uint64 `json:"freeSpace"`
}

type ChunkInfo struct {
    ID   string
    Size uint32
}

type ChunkLocationInfo struct {
    ChunkID      string
    DataNodeIDs  []DataNodeInfo
}