package file

import (
	"encoding/json"
	"hildafetch/internal/utils"
	"os"
	"path"
	"path/filepath"
)

type FileManager struct {
	dir string
}

func NewFileManager(dir string) *FileManager {
	return &FileManager{
		dir: dir,
	}
}

/*
func (f *FileManager) WriteFileDataIfNotExist(files []string) {
			{
		    "username_color": "#80B9B9",
		    "colorPrimary": "#3B6E8F",
		    "colorSecundary": "#D46A6A",
		    "borderColor": "#D46A6A",
		    "borderType": "rounded",

		    "frases": [
		        "Well, that was traumatic... but that’s the life of an adventurer.",
		        "The wilderness is calling.",
		        "Woffs are just misunderstood."
		    ]
		}
	content_template := ">  *OS*: {{OS}}\n>  *PROCESSADOR* : {{CPU}}\n> *Memoria* : {{Memory usage}} / {{Memory}} GiB\n> *Armazenamento* : {{Disk Used}} / {{Disk}} GiB\n> *Tempo de boot* : {{Uptime}}"
	content_config := map[string]interface{}{
		"username_color": "#80B9B9",
		"colorPrimary":   "#3B6E8F",
		"colorSecundary": "#D46A6A",
		"borderColor":    "#D46A6A",
		"borderType":     "rounded",

		"frases": []string{
			"Well, that was traumatic... but that’s the life of an adventurer.",
			"The wilderness is calling.",
			"Woffs are just misunderstood.",
		},
	}
}
*/

func (f FileManager) CreateFileIfNotExist(filename string) (string, error) {
	errDir := os.MkdirAll(filepath.Dir(f.dir), 0755)

	if errDir != nil {
		return "", errDir
	}
	filepath := path.Join(f.dir, filename)

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return "", err
	}

	defer file.Close()

	return file.Name(), nil
}

func (f FileManager) ReadFiles(filename string) ([]byte, error) {
	data, err := os.ReadFile(path.Join(f.dir, filename))

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f FileManager) OpenConfigJson(file string) (*utils.Config, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return &utils.Config{}, err
	}
	var config utils.Config
	errJson := json.Unmarshal(data, &config)

	if errJson != nil {
		return &utils.Config{}, errJson
	}
	return &config, nil
}
