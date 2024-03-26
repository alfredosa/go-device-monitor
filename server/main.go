package main

import (
	"context"
	pb "github.com/alfredosa/go-sensors-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	// Update this import path
)

type server struct {
	pb.UnimplementedMonitorServiceServer
}

func (s *server) SendUsage(ctx context.Context, in *pb.UsageData) (*pb.UsageResponse, error) {
	log.Printf("Received: CPU Usage: %v, RAM Usage: %v%%", in.CpuUsage, in.RamUsage.UsedPercent)
	// Here you can process or store the usage data
	return &pb.UsageResponse{Message: "Usage Received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMonitorServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
