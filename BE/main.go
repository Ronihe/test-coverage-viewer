package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"fuzzbuzz.com/roni/v1/repo"
	git "github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
)

const baseUrl = "https://github.com/"

func main() {

	http.HandleFunc("/info", getRepoInfo)
	http.ListenAndServe(":8080", nil)

}

type data struct {
	Owner    string
	RepoName string
}

func getRepoInfo(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var data data
	err := decoder.Decode(&data)
	if err != nil {
		logrus.WithError(err).Error("could not decode request")
	}

	owner := data.Owner
	repoName := data.RepoName

	repoInfo := repoForOwnerRepoName(owner, repoName)

	js, err := json.Marshal(repoInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

func repoForOwnerRepoName(owner string, repoName string) repo.RepoInfo {
	dir, err := ioutil.TempDir("./", "test")
	if err != nil {
		logrus.WithError(err).Error("could not create temp directory")
		return repo.RepoInfo{}
	}
	defer os.RemoveAll(dir) // clean up after testing coverage is done
	fmt.Println("directory created: ", dir)

	url := fmt.Sprintf("%s%s/%s.git", baseUrl, owner, repoName)
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		logrus.WithError(err).Error("could not clone repo")
		return repo.RepoInfo{}
	}

	repo := repo.GetRepoInfo(repo.CreateNewRepo(owner, repoName, dir))
	fmt.Print(repo.StarNum)
	return repo

}
