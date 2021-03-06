package cli

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var pathJson string = "vscodeCliPath.json"

type pathItem struct {
	Path 		string 		`json:"path"`
	DefaultPath bool		`json:"defaultPath"`
}

type pathContainer struct {
	pathItems []*pathItem
}

var paths *pathContainer

func isExistPathFile() bool {
	_, err := os.Open(pathJson)
	return !os.IsNotExist(err)
}

func (p *pathContainer) changeDefaultPath(path string) {
	for _, pathItem := range p.pathItems {
		if pathItem.Path != path {
			pathItem.DefaultPath = false
		} else {
			pathItem.DefaultPath = true
		}
	} 
}

func (p *pathContainer) getDefaultPath() (defaultPath string) {
	for _, pathItem := range p.pathItems {
		if pathItem.DefaultPath {
			defaultPath = pathItem.Path
			break;
		}
	}
	return 
}

func (p *pathContainer) GetPaths() (allPath []string) {
	for _, pathItem := range p.pathItems {
		allPath = append(allPath, pathItem.Path)
	}
	return
}

func Path() *pathContainer {
	return paths
}

func (p *pathContainer) savePathJson() {
	file, _ := os.Create(pathJson)
	err := json.NewEncoder(file).Encode(p.pathItems)
	defer file.Close()
	if err != nil {
		log.Fatal(errors.New("DON'T SAVE JSON"))
	}
}

func (p *pathContainer) addPath(path string, defaultValue bool) {
	pItem := &pathItem{
		Path: path,
		DefaultPath: defaultValue,
	}
	p.pathItems = append(p.pathItems, pItem)
}

func (p *pathContainer) removePath(path string) error {
	for i, pathItem := range p.pathItems {
		if pathItem.Path == path {
			if pathItem.DefaultPath {
				return errors.New("DEFAULT PATH")
			}
			p.pathItems = append(p.pathItems[:i], p.pathItems[i+1:]...)
			break;
		}
	}
	return nil
}

func InitialPaths() {
	paths = &pathContainer{}
	if !isExistPathFile() {
		data := initialPathSelectPrompt()
		paths.addPath(data, true)
	} else {
		file, _ := os.Open(pathJson)
		json.NewDecoder(file).Decode(&paths.pathItems)
		defer file.Close()
	}	
}
