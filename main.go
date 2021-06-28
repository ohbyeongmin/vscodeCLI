package main

import "github.com/ohbyeongmin/vscodeCLI/cli"


type pathItem struct {
	Path 		string 		`json:"path"`
	DefaultPath bool		`json:"defaultPath"`
}

func main(){
	cli.Start()
	// a := "*obm*"
	// find := exec.Command("find", "/Users/byeongminoh/Documents", "-maxdepth", "1", "-name", a, "-type", "d")
	// res, err := find.Output()
	// fmt.Println(string(res), err)
}