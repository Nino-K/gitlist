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
		Verify(repo, testRepos[i], t)
	}
}

func Verify(repo Repository, testRepo Repository, t *testing.T) {
	if repo.Id != testRepo.Id {
		t.Errorf("Expected a different Id but got :%d", testRepo.Id)
	}
	if repo.Name != testRepo.Name {
		t.Errorf("Expected a different name but got :%s", testRepo.Name)
	}
	if repo.FullName != testRepo.FullName {
		t.Errorf("Expected a different fullName but got :%s", testRepo.FullName)
	}
	if repo.Description != testRepo.Description {
		t.Errorf("Expected a different description but got :%s", testRepo.Description)
	}
	if repo.GitURL != testRepo.GitURL {
		t.Errorf("Expected a different gitURL but got :%s", testRepo.GitURL)
	}
	if repo.CloneURL != testRepo.CloneURL {
		t.Errorf("Expected a different cloneURL but got :%s", testRepo.CloneURL)
	}
	if repo.SshURL != testRepo.SshURL {
		t.Errorf("Expected a different sshURL but got :%s", testRepo.SshURL)
	}

}
