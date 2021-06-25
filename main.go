package main

import (
	"github.com/ohbyeongmin/vscodeCLI/cli"
)


type pathItem struct {
	Path 		string 		`json:"path"`
	DefaultPath bool		`json:"defaultPath"`
}

func main(){
	cli.InitialPaths()
	cli.SelectMenuPrompt()
}