package grpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/grpc-server/interceptor"
)

// DefaultOptions is a set of default gRPC server options that include:
// - Insecure credentials (no TLS)
// - A unary interceptor that validates incoming requests
var (
	DefaultOptions = []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptor.Recover,
			interceptor.Logging,
		),
	}
)

// Service represents a gRPC service with its description and handler.
type Service struct {
	ServiceDesc *grpc.ServiceDesc
	Handler     any
}

// Server represents the gRPC server with its listener and server instance.
type Server struct {
	gRPCServer *grpc.Server
	listener   net.Listener
}

// NewGRPCServer creates and returns a new Server instance listening on the specified port.
// It also registers the user service and reflection service to the gRPC server.
func NewGRPCServer(address string, services []Service, opts ...grpc.ServerOption) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	if len(opts) == 0 {
		opts = DefaultOptions
	}

	gRPCServer := grpc.NewServer(opts...)
	reflection.Register(gRPCServer)

	for _, service := range services {
		gRPCServer.RegisterService(service.ServiceDesc, service.Handler)
	}

	return &Server{
		gRPCServer: gRPCServer,
		listener:   listener,
	}, nil
}

// Start runs the gRPC server in a separate goroutine to handle incoming requests.
func (s *Server) Start() {
	go func() {
		// We can ignore error because
		// Serve will return a non-nil error unless Stop or GracefulStop is called.
		_ = s.gRPCServer.Serve(s.listener)
	}()
}

// Stop gracefully stops the gRPC server, ensuring all ongoing requests are completed.
func (s *Server) Stop() {
	s.gRPCServer.GracefulStop()
}
