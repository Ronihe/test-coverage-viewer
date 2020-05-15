package main

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/sirupsen/logrus"

	git "github.com/go-git/go-git/v5"
)

// /api/:github_org/:github_repo/info
// /api/:github_org/:github_repo/test
const baseUrl = "https://github.com/"

type githubRepo struct {
	Owner string
	Repo  string
}

func main() {
	fmt.Println("hello")
	cloneRepo("google", "uuid")

}

func cloneRepo(owner string, repo string) {
	dir, err := ioutil.TempDir("./", "test")
	if err != nil {
		logrus.WithError(err).Fatal("could not create temp directory")
	}
	defer os.RemoveAll(dir) // clean up after testing coverage is done
	// FIXME: to delete
	fmt.Println("directory created: ", dir)

	url := fmt.Sprintf("%s%s/%s.git", baseUrl, owner, repo)
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		logrus.WithError(err).Fatal("could not clone repo")
	}
	// repo.test()
	// repo.info()
}
