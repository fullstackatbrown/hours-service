package main

import (
	"flag"
	pb "github.com/fullstackatbrown/hours/pb/out"
	"github.com/fullstackatbrown/hours/service"
	"github.com/nthnluu/aether/pkg/server"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func main() {
	flag.Parse()

	// Create the service
	s := service.CreateService(service.CreateInMemoryRepository())

	// Create the gRPC server and register your service.
	grpcServer := server.CreateServer()
	pb.RegisterHoursServiceServer(grpcServer, s)

	// Run the server.
	server.RunOrExit(grpcServer, *port, func(b *server.ServerConfig) {
	})
}
