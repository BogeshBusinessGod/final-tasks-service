package task

import (
	"context"
	"errors"
	"final/interceptors"
	tsk1 "final/pkg/proto/sync/final-boss/v1"
	"fmt"
	"net"
	"net/http"

	"final/internal/config"
	"final/internal/service"
	log "final/internal/utils/observability"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func registerSwagger(mux *http.ServeMux) {
	mux.Handle(swaggerUIRoutePrefix,
		http.StripPrefix(swaggerUIRoutePrefix,
			http.FileServer(http.Dir(swaggerUIFSDir)),
		),
	)

	mux.Handle(swaggerJSONRoutePrefix,
		http.StripPrefix("/swagger",
			http.FileServer(http.Dir(swaggerJSONFSDir)),
		),
	)
}

type Server struct {
	cfg *config.Config
	tsk1.UnimplementedTasksServer
	svc        service.Service
	logger     *log.Logger
	grpcServer *grpc.Server
	httpServer *http.Server
	listener   net.Listener
}

func NewServer(cfg *config.Config, logger *log.Logger, svc service.Service) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		svc:    svc,
	}
}

func (s *Server) Listen() error {
	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.LoggingInterceptor(s.logger),
			interceptors.WithValidation(s.logger),
		),
	)

	tsk1.RegisterTasksServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	gwMux := runtime.NewServeMux()
	if err := tsk1.RegisterTasksHandlerServer(context.Background(), gwMux, s); err != nil {
		return fmt.Errorf("gateway register error: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwMux)

	registerSwagger(mux)

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.cfg.HTTP.Port),
		Handler: mux,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GRPC.Port))
	if err != nil {
		return err
	}
	s.listener = lis

	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			s.logger.Error("gRPC serve failed", err)
		}
	}()

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("HTTP gateway failed", err)
		}
	}()

	s.logger.Info("servers started", "grpc_port", s.cfg.GRPC.Port, "http_port", s.cfg.HTTP.Port)
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping servers...")
	s.grpcServer.GracefulStop()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("HTTP shutdown failed", err)
		return err
	}
	return nil
}
