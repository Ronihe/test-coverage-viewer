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
	dir      string
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
		logrus.WithError(err).Error("get star http get request failed")
	}

	if res.StatusCode != 200 {
		logrus.Error("get star http get request not success")
		return 0
	}

	var starNum starJson
	json.NewDecoder(res.Body).Decode(&starNum)
	return starNum.StarNum
}

func (r *NewRepo) Files() []File {
	fileNameList := getFiles(r.owner, r.repoName, r.dir)
	return fileNameList

}

func getFiles(owner string, repoName string, dir string) []File {
	contentUrl := fmt.Sprintf("%s%s/%s/contents", baseUrl, owner, repoName)
	res, err := http.Get(contentUrl)

	if err != nil {
		logrus.WithError(err).Error("get file http request failed")
		return nil
	}

	if res.StatusCode != 200 {
		logrus.Error("get files http get request not success")
		return nil
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

	err = execTest(dir)
	if err != nil {
		logrus.WithError(err).Error("could not go test the github repo")
		return nil
	}

	coverageMap := ParseFile(dir)

	var testedFileList []File
	for _, fileName := range goFiles {
		content := getContentForFileName(owner, repoName, fileName)

		testedFileList = append(testedFileList, File{
			Name:         fileName,
			Content:      content,
			TestCoverage: coverageMap[fileName],
		})
	}
	return testedFileList
}

func getContentForFileName(owner, repoName, fileName string) string {
	contentUrl := fmt.Sprintf("%s%s/%s/contents/%s", baseUrl, owner, repoName, fileName)
	res, err := http.Get(contentUrl)
	if err != nil {
		logrus.WithError(err).Error("get content http request failed")
		return ""
	}

	if res.StatusCode != 200 {
		logrus.Error("get content http request not success")
		return ""
	}

	var fileContent fileContentJson
	json.NewDecoder(res.Body).Decode(&fileContent)

	return decodeContent(fileContent.Content)
}

// TODO: decode the content in frontend
func decodeContent(encodedString string) string {
	decoded, err := b64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		logrus.WithError(err).Error("could not 64 decode")
		return ""
	}

	return string(decoded)
}

func CreateNewRepo(owner, repoName, dir string) Repo {
	return &NewRepo{
		owner:    owner,
		repoName: repoName,
		dir:      dir,
	}
}
