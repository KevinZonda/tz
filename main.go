package main

import (
	"fmt"
	nvidiasmijson "github.com/fffaraz/nvidia-smi-json"
)

func main() {
	o := nvidiasmijson.XmlToObject(nvidiasmijson.RunNvidiaSmi())
	for i, g := range o.GPUS {
		fmt.Println("IDX:", i)
		fmt.Println(g.ProductName)
		fmt.Println(g.FbMemoryUsageUsed, "/", g.FbMemoryUsageTotal)
	}
}
