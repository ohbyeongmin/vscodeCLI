package cli

import (
	"errors"
	"os/exec"
	"strings"
)


func FindDirectoryPrompt(searchDir string) ([]string ,error) {
	grep := exec.Command("grep", searchDir)
	ls := exec.Command("ls", "-t" )//, path)

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

func InitializeApp() {
	
}

func Start(){
	
}