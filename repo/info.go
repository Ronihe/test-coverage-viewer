package repo

// this package is to get a list of info from github repo
// utlize the http call with net/http
import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type NewRepo struct {
	owner    string
	repoName string
}

type starJson struct {
	StarNum int `json:"stargazers_count"`
}

type fileJson struct {
	Name string `json:"name"`
}

type fileContentJson struct {
	Content string `json:"content"`
}

func (r *NewRepo) StarNum() int {
	starUrl := fmt.Sprintf("%s%s/%s", baseUrl, r.owner, r.repoName)
	res, err := http.Get(starUrl)

	if err != nil {
		logrus.WithError(err).Fatal("get star http get request failed")
	}

	if res.StatusCode != 200 {
		logrus.Fatal("get start http get request not success")
	}

	var starNum starJson
	json.NewDecoder(res.Body).Decode(&starNum)
	return starNum.StarNum
}

func (r *NewRepo) Files() []File {
	fileNameList := getFiles(r.owner, r.repoName)

}

func getFiles(owner string, repoName string) []string {
	contentUrl := fmt.Sprintf("%s%s/%s/contents", baseUrl, owner, repoName)
	fmt.Println(contentUrl)
	res, err := http.Get(contentUrl)

	if err != nil {
		logrus.WithError(err).Fatal("")
	}

	if res.StatusCode != 200 {
		logrus.Fatal("get files http get request not success")
	}

	var fileList []fileJson
	json.NewDecoder(res.Body).Decode(&fileList)

	var goFiles []string
	for _, file := range fileList {
		if file.Name[len(file.Name)-3:] != ".go" {
			continue
		}

		if len(file.Name) > 8 {
			if file.Name[len(file.Name)-8:] == "_test.go" {
				continue
			}
		}

		goFiles = append(goFiles, file.Name)
	}
	return goFiles
}

func getContentForFileName(owner, repoName, fileName string) File {
	contentUrl := fmt.Sprintf("%s%s/%s/contents/%s.go", baseUrl, owner, repoName, fileName)
	res, err := http.Get(contentUrl)
	if err != nil {
		logrus.WithError(err).Fatal("get content http request failed")	
	}

	if res.StatusCode != 200 {
		logrus.Fatal("get content http request not success")
	}

	var fileContent fileContentJson
	json.NewDecoder(res.Body).Decode(&fileContent)
	content := fileContent.Content
	
	return File{
		Name: fileName,
		Content: content,
		// TODO: coverage for this file
	}

	
}

func decodeContent(encodedString string) string {
	decoded, err := b64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		logrus.WithError(err).Fatal("could not 64 decode")
	}

	return string(decoded)
}
