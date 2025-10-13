package api

// ---------- Request Models ----------

// RegisterFileRequest is sent by the client to Metadata Server to start upload.
type RegisterFileRequest struct {
	Filename  string `json:"filename"`
	SizeBytes int64  `json:"sizeBytes"`
}


// ---------- Response Models ----------

// RegisterFileResponse defines what metadata server returns to client after register.
type RegisterFileResponse struct {
	FileID       string            `json:"fileId"`
	ChunkSize    int64             `json:"chunkSize"`
	Chunks       []ChunkAssignment `json:"chunks"`
	Replication  int               `json:"replication"`
	UploadToken  string            `json:"uploadToken"`
	ExpiresAt    string            `json:"expiresAt"`
}

// ChunkAssignment tells client where to upload a chunk.
type ChunkAssignment struct {
	Index   int    `json:"index"`
	ChunkID string `json:"chunkId"`
	ServerAddress string `json:"serverAddress"`
	Offset  int64  `json:"offset"`
	Length  int64  `json:"length"`
}