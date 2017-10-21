package dropbox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewClient(oauthToken string) *Client {
	client := &Client{
		oauthToken: oauthToken,
	}
	return client
}

type Client struct {
	oauthToken string
}

func (c *Client) post(url string, data []byte) ([]byte, error) {
	requestURL := fmt.Sprintf("%s%s", dropboxAPI, url)
	client := http.Client{}
	payload := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", requestURL, payload)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.oauthToken))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func (c *Client) postArg(url string, arg string) ([]byte, error) {
	requestURL := fmt.Sprintf("%s%s", dropboxContent, url)
	client := http.Client{}
	req, err := http.NewRequest("POST", requestURL, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.oauthToken))
	req.Header.Add("Dropbox-API-Arg", arg)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func (c *Client) GetFolderContents(folderPath string) (FolderContents, error) {
	contents := FolderContents{}
	params := struct {
		Path string `json:"path"`
	}{
		Path: folderPath,
	}
	output, _ := json.Marshal(&params)
	data, err := c.post("/files/list_folder", output)
	if err != nil {
		return contents, err
	}
	err = json.Unmarshal(data, &contents)
	return contents, err
}

func (c *Client) GetFile(filePath string) (File, error) {
	file := File{}
	params := struct {
		Path string `json:"path"`
	}{
		Path: filePath,
	}
	output, _ := json.Marshal(&params)
	data, err := c.post("/files/get_metadata", output)
	if err != nil {
		return file, err
	}
	err = json.Unmarshal(data, &file)
	if file.Type != "file" {
		return file, errors.New("invalid file path")
	}
	file.client = c
	return file, err
}
