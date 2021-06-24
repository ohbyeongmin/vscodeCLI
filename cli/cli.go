package cli

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

// const path string = "/Users/byeongminoh/Documents"


const vscodeCliPath = "VSCODE_CLI_PATH"
var path string = "/Users/byeongminoh/Documents"


func FindDirectory(searchDir string) ([]string ,error) {
	grep := exec.Command("grep", searchDir)
	ls := exec.Command("ls", "-t", path)

	pipe, _ := ls.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe
	ls.Start()

	res, _ := grep.Output()

	if len(res) <= 0 {
		return nil, errors.New("NOT FOUND DIRECTORY")
	}

	stringRes := string(res[:])

	return strings.Fields(stringRes), nil
}

func SelectMenu(searchDirs []string) (string, error) {
	templates := &promptui.SelectTemplates{
		Label: 	"\U0001F929 {{ . | cyan}}",
		Active: "\U000025B6 {{ . | green }}",
	}

	items := []string{"CHANGE PATH"}
	items = append(items, searchDirs...)

	prompt := promptui.Select{
		Label: "Select Directory",
		Items: items,
		Size: 10,
		Templates: templates,
	}
	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func setEnvPath() {
	
}

func InitializeConfig() {
	echoPath := exec.Command("sh", "-c", "echo $" + vscodeCliPath)
	bPath, _ := echoPath.Output()
	
	if len(bPath) <= 0 {
		setEnvPath()
	}
}

func Start(){
	// InitializeConfig()
	setEnvPath()
}