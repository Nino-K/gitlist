package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Nino-K/gitlist/apihandler"
	"github.com/Nino-K/gitlist/storage"
	"github.com/olekukonko/tablewriter"
)

var (
	storagePath = fmt.Sprintf("%s/.gitlist", os.Getenv("HOME"))
	strg        = storage.New(storagePath)
)

func queryAndOutput(repoName string) {
	apiUrl := fmt.Sprintf("https://api.github.com/search/repositories?q=%s+language:go&sort=starts", repoName)

	apiHandler := apihandler.New(apiUrl)
	repos, err := apiHandler.GetRepos()
	if err != nil {
		log.Fatalf("something went wrong calling github API: %v", err)
	}

	tableData := storage.ConvertToStorageData(repos)
	outputTable(tableData)
	if err := strg.Encode(tableData); err != nil {
		log.Fatalf("could not catch data: %v", err)
	}
}

func outputTable(data []storage.Data) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"#", "Name", "URL"})
	for _, v := range data {
		repoName := fmt.Sprintf("github.com/%s", v.Repo.FullName)
		table.Append([]string{strconv.Itoa(v.Id), v.Repo.Name, repoName})
	}
	table.Render()
}

func showDetail(id string) {
	dataId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("Error Converting Id: %v", err)
	}
	storage := storage.New(storagePath)
	data, err := storage.GetDataById(dataId)
	if err != nil {
		log.Fatalf("FindById: %v", err)
	}
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{
		"Name",
		"Description",
		"GitURL",
		"CloneURL",
		"SshURL"})

	table.Append([]string{
		data.Repo.Name,
		data.Repo.Description,
		data.Repo.GitURL,
		data.Repo.CloneURL,
		data.Repo.SshURL})
	table.Render()

}
