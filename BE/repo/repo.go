package repo

type Repo interface {
	StarNum() int
	Files() []File
}

// FIXME: should be exported, in the model file
type RepoInfo struct {
	StarNum int    `json:"starNum"`
	GoFiles []File `json:"goFiles"`
}

type File struct {
	Name         string          `json:"name"`
	Content      string          `json:"content"`
	TestCoverage []CoverageBlock `json:"testCoverage"`
}

const baseUrl = "https://api.github.com/repos/"

func GetRepoInfo(repo Repo) RepoInfo {
	starNum := repo.StarNum()
	files := repo.Files()

	return RepoInfo{
		StarNum: starNum,
		GoFiles: files,
	}
}
