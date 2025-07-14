package system

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

type Cpu struct{}

func NewCpuModel() *Cpu {
	return &Cpu{}
}

func (c *Cpu) GetCpuName() (string, error) {
	info, err := cpu.Info()

	if len(info) > 0 {
		return info[0].ModelName, nil
	}

	if err != nil {
		return "", err
	}

	return "", nil

}

func (c *Cpu) GetBootTime() (time.Time, error) {
	boottime, err := host.BootTime()

	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(int64(boottime), 0), nil
}

func (c *Cpu) GetOs() (*host.InfoStat, error) {
	platform, err := host.Info()

	if err != nil {
		return nil, err
	}

	return platform, nil
}
