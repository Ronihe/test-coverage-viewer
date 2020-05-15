package repo

type Repo interface {
	StarNum() int
	Files() []File
}

// FIXME: should be exported, in the model file
type RepoInfo struct {
	StarNum int
	GoFiles []File
}

type File struct {
	Name         string
	Content      string
	TestCoverage string
}

const baseUrl = "https://api.github.com/repos/"

func GetRepoInfo() {
	// init repo struct
	// add to the final

}
