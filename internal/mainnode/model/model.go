package model

type ChunkInfo struct {
    ID   string
    Size uint32
}

type ChunkLocationInfo struct {
    ChunkID      string
    DataNodeIDs  []string
}