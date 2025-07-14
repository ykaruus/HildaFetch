package system

import "github.com/shirou/gopsutil/v3/disk"

type DiskModel struct {
	disk *disk.UsageStat
}

func NewDiskModel(disk *disk.UsageStat) *DiskModel {
	return &DiskModel{
		disk: disk,
	}
}

func (d *DiskModel) GetFreeDisk() uint64 {
	return d.disk.Free
}

func (d *DiskModel) GetUsedDisk() uint64 {
	return d.disk.Used
}

func (d *DiskModel) GetTotalDisk() uint64 {
	return d.disk.Total
}
