package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/PierreBou91/stoRPC/storpc"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type storpcServer struct {
	pb.UnimplementedStorpcServer

	sync.RWMutex
	store map[string]string
}

func (s *storpcServer) PutValue(ctx context.Context, in *pb.Pair) (*pb.PutResponse, error) {
	s.Lock()
	defer s.Unlock()
	s.store[in.Key] = in.Value

	log.Printf("PutValue: %s -> %s", in.Key, in.Value)

	return &pb.PutResponse{
		Ok: true,
	}, nil
}

func (s *storpcServer) GetValue(ctx context.Context, in *pb.Key) (*pb.GetResponse, error) {
	s.RLock()
	defer s.RUnlock()
	value, ok := s.store[in.Key]

	// TODO: More explicit error handling
	if !ok {
		return &pb.GetResponse{
			Ok: false,
		}, nil
	}

	return &pb.GetResponse{
		Ok:    true,
		Value: value,
	}, nil
}

func (s *storpcServer) DeleteValue(ctx context.Context, in *pb.Key) (*pb.DelResponse, error) {
	s.Lock()
	defer s.Unlock()
	delete(s.store, in.Key)
	return &pb.DelResponse{
		Ok: true,
	}, nil
}

func newServer() *storpcServer {
	return &storpcServer{
		store: make(map[string]string),
	}
}

func main() {
	launchServer(":7070")
}

func launchServer(host string) {
	ln, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	healthServer := health.NewServer()

	// There should be logic here to set the status to NOT_SERVING
	healthServer.SetServingStatus(pb.Storpc_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	log.Printf("Server listening on %s", host)

	pb.RegisterStorpcServer(s, newServer())
	healthpb.RegisterHealthServer(s, healthServer)

	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
