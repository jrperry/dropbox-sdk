package dropbox

import "time"

type FolderContents struct {
	Objects []Object `json:"entries"`
	Cursor  string   `json:"cursor"`
	HasMore bool     `json:"has_more"`
}

type Object struct {
	Type       string    `json:".tag"`
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path_lower"`
	Size       int       `json:"size"`
	Hash       string    `json:"content_hash"`
	Rev        string    `json:"rev"`
	ModifiedAt time.Time `json:"server_modified"`
}
