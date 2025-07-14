package main

import (
	"fmt"
	"hildafetch/internal/system"
	"hildafetch/internal/ui"
	"hildafetch/internal/utils"
	"os"
	"os/user"

	"github.com/shirou/gopsutil/v3/disk"
)

func main() {
	//word := []byte{}
	file, err := os.ReadFile("./template/template.txt")

	if err != nil {
		fmt.Println("Não foi Possivel carregar o arquivo de template.txt")
		return
	}
	template := string(file)

	//fmt.Println(string(word))
	config, configErr := utils.OpenConfigJson("./template/config.json")

	if configErr != nil {
		fmt.Println("Não foi Possivel carregar o arquivo de config.json")
		return
	}

	layout := ui.NewLayoutService(config)

	ascii, errFile := ui.NewAsciiService("./template/ascii.txt", layout)

	if errFile != nil {
		fmt.Println("Não foi Possivel carregar o arquivo de ascii.txt")
		return
	}

	disk, _ := disk.Usage("/")

	info := ui.NewInfoService(
		system.NewCpuModel(),
		system.NewMemModel(),
		system.NewDiskModel(disk),
		layout,
	)

	osinfo := info.GetInfo()
	os_username, _ := user.Current()
	result := info.SetInfo(osinfo, template)
	banner := ascii.SetAscii()
	table := layout.CreateTable(result, os_username.Name+"@"+os_username.Username)

	fmt.Println(layout.CreateFetch(banner, table))

}
