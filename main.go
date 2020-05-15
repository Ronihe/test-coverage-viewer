package main

import (
	"fmt"
	"os"

	"io/ioutil"

	"fuzzbuzz.com/roni/v1/repo"

	"github.com/sirupsen/logrus"

	git "github.com/go-git/go-git/v5"
)

// /api/:github_org/:github_repo/info
// /api/:github_org/:github_repo/test
const baseUrl = "https://github.com/"

func main() {
	cloneRepo("google", "uuid")

}

func cloneRepo(owner string, repoName string) {
	dir, err := ioutil.TempDir("./", "test")
	if err != nil {
		logrus.WithError(err).Fatal("could not create temp directory")
	}
	defer os.RemoveAll(dir) // clean up after testing coverage is done
	fmt.Println("directory created: ", dir)

	url := fmt.Sprintf("%s%s/%s.git", baseUrl, owner, repoName)
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		logrus.WithError(err).Fatal("could not clone repo")
	}

	repo := repo.GetRepoInfo(repo.CreateNewRepo(owner, repoName, dir))
	fmt.Print(repo.StarNum)

}
