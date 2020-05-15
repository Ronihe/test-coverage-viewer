package repo

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"io/ioutil"
	"log"
	"os/exec"
)

type ProjectCoverage struct {
	Files []FileCoverage `json:"files"`
}

type FileCoverage struct {
	FileName       string          `json:"file_name"`
	CoverageBlocks []CoverageBlock `json:"coverage_blocks"`
}

type CoverageBlock struct {
	StartLine int `json:"start_line"`
	EndLine   int `json:"end_line"`
}

func execTest(dir string) error {
	cmd := exec.Command("go", "test", "-coverprofile", "coverage.out")
	cmd.Dir = fmt.Sprintf("../%s", dir)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return err
}

func ParseFile(dir string) ProjectCoverage {
	path := fmt.Sprintf("../%s/coverage.out")
	input, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.WithError(err).Fatal("could not read coverage.out")
	}
	coverage := ParseCover(input)

	for _, file := range coverage.Files {
		fmt.Println(file)
	}
	return coverage
}

func ParseCover(coverFile []byte) ProjectCoverage {
	var files []FileCoverage
	// Go Cover File returns code blocks categorized by file, so we can
	// just have one current file
	currentFile := FileCoverage{}
	for _, line := range bytes.Split(coverFile, []byte("\n")) {
		if bytes.Contains(line, []byte("mode: ")) || len(line) == 0 {
			continue
		}
		// splits into lines formatted like:
		// <package_name>/<file_name>.go:<line_start>.<col_start>,<line_end>.<col_end> <numstatements> <count>
		// We want package/file name, line_start and line_end
		firstPart := string(bytes.Split(line, []byte(" "))[0])
		colonSplit := strings.Split(firstPart, ":")
		fileName := colonSplit[0]
		if !strings.HasSuffix(fileName, ".go") {
			// This isn't a well-formatted line
			continue
		}

		lineNumbersSplit := strings.Split(colonSplit[1], ",")
		startLine, _ := strconv.Atoi(strings.Split(lineNumbersSplit[0], ".")[0])
		endLine, _ := strconv.Atoi(strings.Split(lineNumbersSplit[1], ".")[0])

		if fileName != currentFile.FileName {
			if len(currentFile.CoverageBlocks) > 0 {
				files = append(files, currentFile)
			}
			currentFile = FileCoverage{
				FileName: fileName,
				CoverageBlocks: []CoverageBlock{{
					StartLine: startLine,
					EndLine:   endLine,
				}},
			}
		} else {
			currentFile.CoverageBlocks = append(currentFile.CoverageBlocks, CoverageBlock{
				StartLine: startLine,
				EndLine:   endLine,
			})
		}
	}

	return ProjectCoverage{
		Files: append(files, currentFile),
	}
}
