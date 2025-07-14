package ui

import (
	"fmt"
	"hildafetch/internal/system"
	"strings"
	"time"
)

type InfoItem struct {
	label string
	value string
}

type InfoService struct {
	cpu  *system.Cpu
	mem  *system.MemModel
	disk *system.DiskModel

	layout *LayoutService
}

func NewInfoService(
	cpu *system.Cpu,
	mem *system.MemModel,
	disk *system.DiskModel,
	layout *LayoutService) *InfoService {
	return &InfoService{
		cpu:    cpu,
		mem:    mem,
		disk:   disk,
		layout: layout,
	}
}

func (info *InfoService) GetInfo() map[string]string {
	cpuName, cpuErr := info.cpu.GetCpuName()
	mem, _ := info.mem.GetMem()
	memUsed := float64(mem["memory_used"]) / 1024 / 1024
	memFree := float64(mem["memory_free"]) / 1024 / 1024
	memTotal := float64(mem["memory_total"]) / 1024 / 1024
	memUsedPercent := float64(mem["memory_used_percent"])
	boottime, _ := info.cpu.GetBootTime()
	platform, platformErr := info.cpu.GetOs()
	uptime := time.Since(boottime)

	diskFree := float64(info.disk.GetFreeDisk()) / 1024 / 1024
	diskTotal := float64(info.disk.GetTotalDisk()) / 1024 / 1024
	diskUsed := float64(info.disk.GetUsedDisk()) / 1024 / 1024

	fmrt_bootime := strings.Split(boottime.String(), " ")

	if cpuErr != nil {
		cpuName = "Null"
	}

	if platformErr != nil {
		platform = nil
	}

	return map[string]string{
		"OS":                  platform.Platform + " " + platform.PlatformVersion + " " + platform.PlatformFamily,
		"CPU":                 cpuName,
		"Memory usage":        fmt.Sprintf("%.2f", memUsed),
		"Memory Free":         fmt.Sprint(memFree),
		"Memory":              fmt.Sprintf("%.2f", memTotal),
		"Memory Used Percent": fmt.Sprint(float64(memUsedPercent)),
		"Disk Free":           fmt.Sprintf("%.2f", diskFree),
		"Disk":                fmt.Sprintf("%.2f", diskTotal),
		"Disk Used":           fmt.Sprintf("%.2f", diskUsed),
		"BootTime":            fmrt_bootime[0] + " " + fmrt_bootime[1],
		"Uptime":              fmt.Sprint(uptime.Round(time.Second)),
	}

}

func (info *InfoService) SetInfo(os map[string]string, template string) string {
	for key, value := range os {
		placeholder := "{{" + key + "}}"
		placeholderUmp := key + "&"

		result := info.layout.CreateColumText(key)

		template = strings.ReplaceAll(template, placeholder, value)
		template = strings.ReplaceAll(template, placeholderUmp, result)

	}

	return template
}
