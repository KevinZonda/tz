package handler

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/KevinZonda/tz/shared"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func cpuInfo() string {
	var result strings.Builder
	percentages, err := cpu.Percent(0, true)
	if err != nil {
		log.Println(err)
		return "Error getting CPU percentages\n"
	}
	result.WriteString(fmt.Sprintf("Total Cores: %d", runtime.NumCPU()))
	result.WriteString("    ")

	load := 0.0
	for _, percentage := range percentages {
		load += percentage
	}

	result.WriteString(fmt.Sprintf("Load: %.2f%%", load))
	result.WriteString("    ")
	result.WriteString(fmt.Sprintf("Avg Load (Per Core): %.2f%%", load/float64(len(percentages))))
	result.WriteString("\n\n")

	// If there are many cores, display in a grid
	if len(percentages) <= 8 {
		for i, percentage := range percentages {
			result.WriteString(fmt.Sprintf("Core #%d: %.2f%%\n", i, percentage))
		}
		return result.String()
	}

	cols := 4
	for i, percentage := range percentages {
		if i > 0 && i%cols == 0 {
			result.WriteString("\n")
		}
		result.WriteString(fmt.Sprintf("Core %-2d: %6.2f%%    ", i, percentage))
	}
	result.WriteString("\n")
	return result.String()
}

func memInfo() string {
	var result strings.Builder
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		result.WriteString(fmt.Sprintf("Error getting memory info: %v\n", err))
	}
	result.WriteString(fmt.Sprintf("Total: %.2f GB", float64(memInfo.Total)/(1024*1024*1024)))
	result.WriteString("\n")
	result.WriteString(fmt.Sprintf("Used: %.2f GB (%.2f%%)",
		float64(memInfo.Used)/(1024*1024*1024),
		memInfo.UsedPercent))
	result.WriteString("\n")
	result.WriteString(fmt.Sprintf("Free: %.2f GB", float64(memInfo.Free)/(1024*1024*1024)))
	result.WriteString("\n")
	return result.String()
}

func Status() string {
	var result strings.Builder
	result.WriteString(wrap("CPU"))
	result.WriteString(cpuInfo())

	result.WriteString("\n")
	result.WriteString(wrap("MEM"))
	result.WriteString(memInfo())

	result.WriteString("\n")
	result.WriteString(wrap("GPU"))
	if shared.IsNvAvailable {
		result.WriteString(shared.GetSmi())
	} else {
		result.WriteString("NV SMI is not available.\n")
	}
	return result.String()
}

func wrap(s string) string {
	return fmt.Sprintf(`#######
# %s #
#######

`, s)
}

func Cpu() string {
	return cpuInfo()
}

func Mem() string {
	return memInfo()
}
