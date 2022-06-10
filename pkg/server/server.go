package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/rotationalio/agenda/pkg"
	"github.com/rotationalio/agenda/pkg/api/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	api.UnimplementedAgendaServer
	srv   *grpc.Server
	echan chan error
}

func New() (*Server, error) {
	s := &Server{
		echan: make(chan error, 1),
	}

	// Create the gRPC Server
	s.srv = grpc.NewServer()
	api.RegisterAgendaServer(s.srv, s)
	return s, nil
}

func (s *Server) Serve(addr string) (err error) {
	// Catch OS signals for graceful shutdowns
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		s.echan <- s.Shutdown()
	}()

	// Listen for TCP requests on the specified address and port
	var sock net.Listener
	if sock, err = net.Listen("tcp", addr); err != nil {
		return fmt.Errorf("could not listen on %q", addr)
	}

	// Run the server
	go s.Run(sock)
	log.Info().Str("listen", addr).Str("version", pkg.Version()).Msg("agenda server started")

	// Block and wait for an error either from shutdown or grpc.
	if err = <-s.echan; err != nil {
		return err
	}
	return nil
}

func (s *Server) Run(sock net.Listener) {
	defer sock.Close()
	if err := s.srv.Serve(sock); err != nil {
		s.echan <- err
	}
}

func (s *Server) Shutdown() error {
	log.Info().Msg("server shutting down")
	s.srv.GracefulStop()
	log.Debug().Msg("server has shut down gracefully")
	return nil
}

func (s *Server) Daily(ctx context.Context, in *api.Day) (out *api.Docket, err error) {
	return nil, status.Error(codes.Unimplemented, "the daily rpc is not yet implemented")
}

func (s *Server) Schedule(ctx context.Context, in *api.Item) (out *api.Item, err error) {
	return nil, status.Error(codes.Unimplemented, "the schedule rpc is not yet implemented")
}
