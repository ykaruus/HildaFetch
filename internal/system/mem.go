package system

import "github.com/shirou/gopsutil/v3/mem"

type MemModel struct{}

func NewMemModel() *MemModel {
	return &MemModel{}
}

func (m *MemModel) GetMem() (map[string]uint64, error) {
	mem, err := mem.VirtualMemory()

	if err != nil {
		return map[string]uint64{}, err
	}

	return map[string]uint64{
		"memory_free":          mem.Free,
		"memory_total":         mem.Total,
		"memory_used":          mem.Used,
		"memory_used_percent ": uint64(mem.UsedPercent),
	}, nil

}
