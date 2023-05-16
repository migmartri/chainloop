package main

import (
	"context"
	"fmt"

	pb "github.com/chainloop-dev/chainloop/app/controlplane/plugin/api/gen"
	"google.golang.org/grpc"
)

func SayHi(who string) {
	fmt.Println("hello", who)
}

type Server struct {
	pb.UnimplementedPluginServiceServer
}

func InitializeServer(srv *grpc.Server) {
	pb.RegisterPluginServiceServer(srv, &Server{})
}

func (s *Server) Hi(ctx context.Context, req *pb.HiRequest) (*pb.HiResponse, error) {
	return nil, nil
}
