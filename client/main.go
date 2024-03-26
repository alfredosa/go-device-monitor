package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type MemoryUtilization struct {
	total uint64
	free  uint64
	used  float64
}

type CpuUtilization struct {
	times      string
	cpuPercent float64
}

func main() {
	var (
		memUtil MemoryUtilization
		cpuUtil CpuUtilization
		publish bool
	)

	for {
		publish = false
		time.Sleep(3 * time.Second)
		v, _ := mem.VirtualMemory()
		cpuprcnt, _ := cpu.Percent(0, false)

		cpuUtil.cpuPercent = cpuprcnt[0]
		fmt.Printf("CPU percentage: %f%%\n", cpuUtil.cpuPercent)

		if v.UsedPercent != memUtil.used {
			publish = true
		}

		memUtil.free = v.Free
		memUtil.used = v.UsedPercent
		memUtil.total = v.Total

		if publish {
			fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
		} else {
			fmt.Printf("nothing to report\n")
		}
	}
}
