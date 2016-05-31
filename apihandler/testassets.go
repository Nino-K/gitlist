package apihandler

import (
	"fmt"
	"testing"
	"time"
)

func GetTestRepos(repos ...string) GithubResponse {
	testRepos := make([]Repository, len(repos))
	for i, repo := range repos {
		repoName := fmt.Sprintf("%s%d", repo, i)
		githubAcc := time.Now().Unix()
		testRepos[i] = Repository{
			Id:          i,
			Name:        repoName,
			FullName:    "full" + repoName,
			Description: fmt.Sprintf("some description, %d", githubAcc),
			GitURL:      fmt.Sprintf("git://github.com/%d/%s.git", githubAcc, repoName),
			CloneURL:    fmt.Sprintf("https://github.com/%d/%s.git", githubAcc, repoName),
			SshURL:      fmt.Sprintf("git@github.com:%d/%s.git", githubAcc, repoName),
		}
	}
	return GithubResponse{Items: testRepos}
}

func VerifyItems(repos []Repository, testRepos []Repository, t *testing.T) {
	for i, repo := range repos {
		if repo.Id != testRepos[i].Id {
			t.Errorf("Expected a different Id but got :%d", testRepos[i].Id)
		}
		if repo.Name != testRepos[i].Name {
			t.Errorf("Expected a different name but got :%s", testRepos[i].Name)
		}
		if repo.FullName != testRepos[i].FullName {
			t.Errorf("Expected a different fullName but got :%s", testRepos[i].FullName)
		}
		if repo.Description != testRepos[i].Description {
			t.Errorf("Expected a different description but got :%s", testRepos[i].Description)
		}
		if repo.GitURL != testRepos[i].GitURL {
			t.Errorf("Expected a different gitURL but got :%s", testRepos[i].GitURL)
		}
		if repo.CloneURL != testRepos[i].CloneURL {
			t.Errorf("Expected a different cloneURL but got :%s", testRepos[i].CloneURL)
		}
		if repo.SshURL != testRepos[i].SshURL {
			t.Errorf("Expected a different sshURL but got :%s", testRepos[i].SshURL)
		}
	}

}
