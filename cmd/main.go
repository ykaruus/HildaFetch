package main

import (
	"fmt"
	"hildafetch/internal/file"
	"hildafetch/internal/system"
	"hildafetch/internal/ui"
	"hildafetch/internal/utils"
	"os/user"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	const dir = "template/"
	files := []string{
		"config.json",
		"template.txt",
		"ascii.txt",
	}

	fileManager := file.NewFileManager(dir)

	for _, file := range files {
		if filename, err := fileManager.CreateFileIfNotExist(file); err != nil {
			fmt.Printf("Não foi possivel criar o arquivo : %s\n ", filename)
			return
		}
	}

	config, configErr := fileManager.OpenConfigJson(files[0])

	if configErr != nil {
		utils.SetConfigDefaults(config)
		fmt.Println("Não foi Possivel carregar o arquivo de config.json\nSetando configurações padrão")
	}

	template, err := fileManager.ReadFiles(files[1])
	templateString := string(template)
	if err != nil || len(template) <= 0 {
		templateString = utils.SetTemplateDefaults()
	}

	asciiData, asciiErr := fileManager.ReadFiles(files[2])

	if asciiErr != nil {
		return
	}

	layout := ui.NewLayoutService(config)

	ascii := ui.NewAsciiService(asciiData, layout)

	partition, _ := disk.Usage("/")

	info := ui.NewInfoService(
		system.NewCpuModel(),
		system.NewMemModel(),
		system.NewDiskModel(partition),
		layout,
	)

	osinfo := info.GetInfo()
	os_username, _ := user.Current()
	result := info.SetInfo(osinfo, templateString)
	banner := ascii.SetAscii()
	table := layout.CreateTable(result, os_username.Name+"@"+os_username.Username)

	fmt.Println(layout.CreateFetch(banner, table))

}
