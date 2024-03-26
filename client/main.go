package main

import (
	"context"
	"fmt"
	pb "github.com/alfredosa/go-sensors-client/proto"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"google.golang.org/grpc"
	"log"
	"time"
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

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMonitorServiceClient(conn)

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

		if v.UsedPercent != memUtil.used {
			publish = true
		}

		memUtil.free = v.Free
		memUtil.used = v.UsedPercent
		memUtil.total = v.Total

		if publish {
			fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := c.SendUsage(ctx, &pb.UsageData{
				CpuUsage: float32(cpuUtil.cpuPercent),
				RamUsage: &pb.MemoryUsage{
					Free:        float32(memUtil.free),
					UsedPercent: float32(memUtil.used),
					Total:       float32(memUtil.total),
				},
			})
			if err != nil {
				log.Fatalf("could not send usage: %v", err)
			}
			log.Printf("Response: %s", r.Message)
		} else {
			fmt.Printf("nothing to report\n")
		}
	}
}
