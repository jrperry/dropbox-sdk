package dropbox

import (
	"encoding/json"
	"time"
)

type File struct {
	client     *Client
	Type       string    `json:".tag"`
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	URL        string    `json:"url"`
	Path       string    `json:"path_lower"`
	Size       int       `json:"size"`
	Hash       string    `json:"content_hash"`
	Rev        string    `json:"rev"`
	ModifiedAt time.Time `json:"server_modified"`
}

func (f *File) Share() error {
	params := struct {
		Path string `json:"path"`
	}{
		Path: f.Path,
	}
	output, _ := json.Marshal(&params)
	data, err := f.client.post("/sharing/create_shared_link_with_settings", output)
	if err != nil {
		return err
	}
	resp := struct {
		URL string `json:"url"`
	}{}
	err = json.Unmarshal(data, &resp)
	f.URL = resp.URL
	return err
}

func (f *File) GetThumbnail() []byte {
	params := struct {
		Path   string `json:"path"`
		Format string `json:"format"`
		Size   string `json:"size"`
	}{
		Path:   f.Path,
		Format: "jpeg",
		Size:   "w128h128",
	}
	output, _ := json.Marshal(&params)
	data, _ := f.client.postArg("/files/get_thumbnail", string(output))
	return data
}
