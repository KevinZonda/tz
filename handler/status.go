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
	result.WriteString(fmt.Sprintf("Cores: %d\n", runtime.NumCPU()))
	for i, percentage := range percentages {
		result.WriteString(fmt.Sprintf("Core #%d: %.2f%%\n", i, percentage))
	}
	return result.String()
}

func memInfo() string {
	var result strings.Builder
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		result.WriteString(fmt.Sprintf("Error getting memory info: %v\n", err))
	}
	result.WriteString(fmt.Sprintf("Total: %.2f GB\n", float64(memInfo.Total)/(1024*1024*1024)))
	result.WriteString(fmt.Sprintf("Used: %.2f GB (%.2f%%)\n",
		float64(memInfo.Used)/(1024*1024*1024),
		memInfo.UsedPercent))
	result.WriteString(fmt.Sprintf("Free: %.2f GB\n", float64(memInfo.Free)/(1024*1024*1024)))
	return result.String()
}

func Status() string {
	var result strings.Builder
	result.WriteString("CPU\n")
	result.WriteString(cpuInfo())

	result.WriteString("\nMEM\n")
	result.WriteString(memInfo())

	result.WriteString("\nGPU\n")
	if shared.IsNvAvailable {
		result.WriteString(shared.GetSmi())
	} else {
		result.WriteString("NV SMI is not available.\n")
	}
	return result.String()
}
