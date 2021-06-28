package cli

import (
	"fmt"
	"os/exec"
	"strings"
)

func execProject(project string) {
	code := exec.Command("code", project)
	code.Run()
}


func findDefaultDirectory(searchDir string) ([]string) {
	var foundDirs []string
	path := Path().GetDefaultPath()

	
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

func InitializeApp() {
	
}

func Start(){
	InitialPaths()
	var searchDir string
	var searchDirResults []string
	for {
		switch SelectMenuPrompt() {
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
				Path().ChangeDefaultPath(defaultPath)
			case menu["EXIT"]:
				defer Path().SavePathJson()
				return
		}
	}

}