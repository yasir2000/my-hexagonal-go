package rpc

import (
	"log"
	"net"

	"github.com/yasir2000/my-hexagonal-go/adaptors/framework/left/grpc/pb"
	"github.com/yasir2000/my-hexagonal-go/internal/ports"
	"google.golang.com/grpc"
)

type Adaptor struct {
	api ports.APIPort
}

func NewAdaptor(api ports.APIPort) *Adaptor {
	return &Adaptor{api: api}
}

func (grpca Adaptor) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen to port 9000: %v", err)
		arithmaticServiceServer := grpca
		grpcServer := grpc.NewServer()
		pb.RegisterArithmaticServiceServer(grpcServer, arithmaticServiceServer)
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
		}

	}
}
