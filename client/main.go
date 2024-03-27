package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/alfredosa/go-sensors-client/proto"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"google.golang.org/grpc"
)

// BUG: SHOULD only track and send data. Do not store unecessary stuff on the heap :D it should be lightweight
func main() {
	// TODO: GET THIS REMOVED :) IT was to be read byu variable set in docker
	// TODO: remove grpc with insecure, its deprecated
	server_host := os.Getenv("SERVER_HOST")

	// NOTE: For now we use localhost :D
	if server_host == "" {
		log.Printf("Server host not set")

		server_host = "localhost"
	}

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		log.Fatalf("Server port not set")
	}

	conn, err := grpc.Dial(server_host+":"+port, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMonitorServiceClient(conn)

	for {
		time.Sleep(10 * time.Second)

		v, _ := mem.VirtualMemory()
		memoryUsage := &pb.MemoryUsage{
			Free:        float32(v.Free),
			UsedPercent: float32(v.UsedPercent),
			Total:       float32(v.Total),
		}

		hostDetails, err := getHostInfo()

		if err != nil {
			log.Fatalf("Error getting host information %s", err)
		}

		cpuUsage, err := getCPUUsage()
		if err != nil {
			log.Fatalf("Error getting cpu usage %s", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SendUsage(ctx, &pb.UsageData{
			CpuUsage:    cpuUsage,
			RamUsage:    memoryUsage,
			HostDetails: hostDetails,
		})
		if err != nil {
			log.Fatalf("could not send usage: %v", err)
		}
		log.Printf("Response: %s", r.Message)
	}
}

func getCPUUsage() (float32, error) {

	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		fmt.Printf("Failed getting the cpu percent: %s", err)
	}

	return float32(cpuPercent[0]), nil
}

func getHostInfo() (*pb.HostDetails, error) {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Printf("Failed getting the host: %s", err)
	}

	return &pb.HostDetails{
		Hostname:      hostInfo.Hostname,
		Uptime:        hostInfo.Uptime,
		KernelVersion: hostInfo.KernelVersion,
		KernelArch:    hostInfo.KernelArch,
		Os:            hostInfo.OS,
		BootTime:      hostInfo.BootTime,
		HostId:        hostInfo.HostID,
	}, nil
}
