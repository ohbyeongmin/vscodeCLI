package cli

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)


func DirectExecProject(searchDir string) error{
	var foundDirs []string
	paths := paths.GetPaths()
	

	for _, path := range paths {
		find := exec.Command("find", path, "-maxdepth", "1", "-name", searchDir, "-type", "d")
		res, _ := find.Output()

		if len(res) <= 0 {
			continue
		}
		stringRes := strings.Fields(string(res[:]))
		foundDirs = append(foundDirs, stringRes...)
	}

	if len(foundDirs) == 0 {
		return errors.New("ERROR")
	} else if len(foundDirs) == 1 {
		execProject(foundDirs[0])
	} else {
		s := selectProjectPrompt(foundDirs)
		execProject(s)
	}
	return nil
}


func execProject(project string) {
	code := exec.Command("code", project)
	code.Run()
}


func findDefaultDirectory(searchDir string) ([]string) {
	var foundDirs []string
	path := Path().getDefaultPath()

	
	searchDir = fmt.Sprintf("*%s*",searchDir)
	find := exec.Command("find", path, "-maxdepth", "1", "-name", searchDir, "-type", "d")
	res, _ := find.Output()

	stringRes := strings.Fields(string(res[:]))
	foundDirs = append(foundDirs, stringRes...)
	

	return foundDirs
}

func findAllDirectory(searchDir string) ([]string) {
	var foundDirs []string

	paths := paths.GetPaths()

	for _, path := range paths {
		searchDir = fmt.Sprintf("*%s*",searchDir)
		find := exec.Command("find", path, "-maxdepth", "1", "-name", searchDir, "-type", "d")
		res, _ := find.Output()

		if len(res) <= 0 {
			continue
		}

		stringRes := strings.Fields(string(res[:]))
		foundDirs = append(foundDirs, stringRes...)
	}

	return foundDirs
}

func InitializeApp(){
	InitialPaths()
}

func Start(){
	var searchDir string
	var searchDirResults []string
	for {
		switch selectMenuPrompt() {
			case menu["ALL"]:
				searchDir = searchDirPrompt()
				searchDirResults = findAllDirectory(searchDir)
				selectedDir := selectProjectPrompt(searchDirResults)
				execProject(selectedDir)		
			case menu["DEFAULT"]:
				searchDir = searchDirPrompt()
				searchDirResults = findDefaultDirectory(searchDir)
				selectedDir := selectProjectPrompt(searchDirResults)
				execProject(selectedDir)
			case menu["CHANGE"]:
				defaultPath := changeDefaultPathPrompt()
				Path().changeDefaultPath(defaultPath)
			case menu["ADD"]:
				aPath := addPathPrompt()
				Path().addPath(aPath, false)
			case menu["REMOVE"]:
				rPath := removePathPrompt()
				err := Path().removePath(rPath)
				if err != nil {
					fmt.Println("Default Path는 삭제 할 수 없습니다.")
				}
			case menu["EXIT"]:
				defer Path().savePathJson()
				return
		}
	}

}