package apihandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	STANDARD_HEADER = "application/vnd.github.v3.text-match+json"
)

type Repository struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	GitURL      string `json:"git_url"`
	CloneURL    string `json:"clone_url"`
	SshURL      string `json:"ssh_url"`
}

type GithubResponse struct {
	Items []Repository `json:"items"`
}

type ApiHandler struct {
	url string
}

func New(url string) *ApiHandler {
	return &ApiHandler{url: url}
}

func (a *ApiHandler) GetRepos() ([]Repository, error) {
	httpResponse, err := get(a.url)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	content, err := ioutil.ReadAll(httpResponse.Body)
	var repos GithubResponse
	err = json.Unmarshal(content, &repos)
	if err != nil {
		return nil, err
	}
	return repos.Items, nil
}

func get(url string) (*http.Response, error) {
	if url == "" {
		return nil, &apiHandlerError{url, "url can not be empty"}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", STANDARD_HEADER)

	client := &http.Client{}

	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("something went wrong %s", resp.Status))
	}
	return resp, err
}
