package api

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/viper"
)

const (
	verifyGitRepoUrl = "%s/repos"
	cloneGitRepoUrl = "https://github.com/%s/%s.git"
)

func CheckGitRepoExist(name string) bool {
	newUrl := verifyGitRepoUrl + fmt.Sprintf("/%s/%s", viper.GetString("user"), name)
	response := makeRequest("GET", newUrl, nil)

	return response.error == nil
}

func CloneGitRepoLocal(userName, repoName string) (bool, error) {
	localDir := viper.GetString("GIT_DIR")

	if _, err := os.Stat(localDir); !os.IsNotExist(err) {
		os.RemoveAll(localDir) // clean up
	}

	_, err := git.PlainClone(localDir, false, &git.CloneOptions{
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: userName,
			Password: viper.GetString("token"),
		},
		URL: fmt.Sprintf(cloneGitRepoUrl, userName, repoName),
		Progress: os.Stdout,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}