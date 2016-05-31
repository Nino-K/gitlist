package storage

import (
	"encoding/gob"
	"os"

	"github.com/nino-k/gitlist/apihandler"
)

type Storage struct {
	path string
}

type Data struct {
	Id   int
	Repo apihandler.Repository
}

func New(path string) *Storage {
	return &Storage{path: path}
}

func (s *Storage) Encode(data []Data) error {
	file, err := os.Create(s.path)
	defer file.Close()
	if err != nil {
		return err
	}

	dataEncoder := gob.NewEncoder(file)
	return dataEncoder.Encode(data)
}

func (s *Storage) Decode() ([]Data, error) {
	dataFile, err := os.Open(s.path)
	defer dataFile.Close()
	if err != nil {
		return nil, err
	}

	var decodedData []Data
	decoder := gob.NewDecoder(dataFile)
	err = decoder.Decode(&decodedData)
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}

func ConvertToStorageData(repos []apihandler.Repository) []Data {
	data := make([]Data, len(repos))
	for i, repo := range repos {
		data[i] = Data{Id: i + 1, Repo: repo}
	}
	return data
}

func ConvertToRepository(data []Data) []apihandler.Repository {
	repos := make([]apihandler.Repository, len(data))
	for k, v := range data {
		repos[k] = v.Repo
	}
	return repos
}