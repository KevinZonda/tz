package main

import (
	"fmt"
	nvidiasmijson "github.com/KevinLinuxTool/nvidia-smi-json"
	"strings"
)

func Status() string {
	sb := &strings.Builder{}

	o := nvidiasmijson.XmlToObject(nvidiasmijson.RunNvidiaSmi())
	var gpus []string
	for i, g := range o.GPUS {
		gpuInfoBuf := &strings.Builder{}
		fmt.Fprintf(gpuInfoBuf, "[%d] %s (%s)\n", i, g.ProductName, g.PowerState)
		fmt.Fprintf(gpuInfoBuf, "MEM: %s / %s\n", g.FbMemoryUsageUsed, g.FbMemoryUsageTotal)
		fmt.Fprintf(gpuInfoBuf, "FAN: %s\n", g.FanSpeed)
		fmt.Fprintf(gpuInfoBuf, "TEM: %s (Current) / %s (Slow) / %s (Max)\n", g.GpuTemp, g.GpuTempSlowThreshold, g.GpuTempMaxGpuThreshold)
		fmt.Fprintf(gpuInfoBuf, "PWR: %s / %s\n", g.PowerDraw, g.EnforcedPowerLimit)
		gpuInfoBuf.WriteString("-------------------------------")
		gpus = append(gpus, gpuInfoBuf.String())
	}
	sb.WriteString(strings.Join(gpus, "-------------------------------\n"))
	return sb.String()
}
