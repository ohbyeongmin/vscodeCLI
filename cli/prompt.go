package cli

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)


var templates = &promptui.SelectTemplates{
	Label: 	"\U0001F929 {{ . | cyan}}",
	Active: "\U000025B6 {{ . | green }}",
}

func addPathPrompt() string {
	validate := func(input string) error {
		ls := exec.Command("ls", input)
		_, err := ls.Output()
		if err != nil {
			return errors.New("Invalid Path")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "추가 할 경로를 입력해 주세요",
		Validate: validate,
	}

	result, _ := prompt.Run()

	return result
}

func initialPathSelectPrompt() string {
	var initialPath string
	pwdCmd := exec.Command("pwd")
	bPath, _ := pwdCmd.Output()
	currentPath := strings.Trim(string(bPath[:]), "\r\n")

	items := []string{currentPath + " (현재 경로)", "다른 경로 선택 (직접 입력)"}

	prompt := promptui.Select{
		Label: "Defalt Path 설정",
		Items: items,
		Size: 5,
		Templates: templates,
	}

	_, result, _ := prompt.Run()

	switch result {
		case items[0]:
			initialPath = currentPath
		case items[1]:
			initialPath = addPathPrompt()
	}
	return initialPath
}

func ChangeDefaultPathPrompt(dirList []string) (string, error) {
	items := []string{"CHANGE PATH"}
	items = append(items, dirList...)

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



func SelectMenuPrompt() {
	defaultPath := Path().GetDefaultPath()
	items := []string{"All", "Defalt"}
	prompt := promptui.Select{
		Label: fmt.Sprintf("Select Menu (default: %s)", defaultPath),
		Items: items,
		Size: 10,
		Templates: templates,
	}
	_, result, _ := prompt.Run()
	fmt.Println(result)
}
