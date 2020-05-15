package main

import (
	"fmt"
	"log"

	"context"
	"io/ioutil"

	git "github.com/go-git/go-git/v5"
	github "github.com/google/go-github/v31/github"
)

// /api/:github_org/:github_repo/info
// /api/:github_org/:github_repo/test
var url = "https://github.com/google/uuid.git"

type githubRepo struct {
	Owner string
	Repo  string
}

func main() {
	fmt.Println("hello")
	cloneRepo()
	getInfo()

}

// TODO: should create a package to make the

func cloneRepo() {
	fmt.Println("I am clonding ")
	dir, err := ioutil.TempDir("./", "test")

	if err != nil {
		fmt.Println(err)
	}
	// DELAY: AND DELETE THE REPO
	// TEST AND GO FILE AND GET THE COVERAGE FILE
	// THE JSON FILE AND UTILIZE FUNCTION CONVERT TO THE FILE TO FRONT END
	//

	fmt.Println(dir)
	// Clones the repository into the given dir, just as a normal git clone does
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func removeRepo() {

	fmt.Println("I am removing the ")

}

func testRepo() {
	fmt.Println("testing")
	// use the exec commnad to executee the go test, returnt he file
}

func parse2Json() {
	fmt.Println("parsing json--")
	// can be pase to json from the file
	// send to the front end

}

func getInfo() {
	// list of files
	// stars
	client := github.NewClient(nil)

	_, resp, err := client.Activity.ListStargazers(context.TODO(), "google", "uuid", nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*resp)
}
