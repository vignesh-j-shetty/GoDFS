package api

// ---------- Request Models ----------
type FolderRequest struct {
	Path       string `json:"path"`
	FolderName string `json:"folderName"`
}
