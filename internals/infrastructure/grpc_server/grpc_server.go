package grpcserver

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller"
	controller_process_tracing "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/product"

	// "git.robodev.co/imp/shared-utility/grpc_errors"
	// "git.robodev.co/imp/shared-utility/validator"
	api_v1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
	grpc_health_v1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/grpc/health/v1"

	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	Config                   config.Configuration
	Server                   *grpc.Server
	HealthCtrl               *controller.HealthZController
	PingPongCtrl             *controller.PingPongController
	ProcessTracingController *controller_process_tracing.Controller
}

// Configure ...
func (s *Server) Configure() {
	grpc_health_v1.RegisterHealthServer(s.Server, s.HealthCtrl)
	api_v1.RegisterPingPongServiceServer(s.Server, s.PingPongCtrl)
	api_v1.RegisterProductServiceServer(s.Server, s.ProcessTracingController)
}

// Start ...
func (s *Server) Start() {
	go func() {
		listen, err := net.Listen("tcp", ":"+strconv.Itoa(s.Config.Port))

		if err != nil {
			panic(err)
		}

		if err := s.Server.Serve(listen); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Listening and serving HTTP on", strconv.Itoa(s.Config.Port))

	// Gracefully Shutdown
	// Make channel listen for signals from OS
	gracefulStop := make(chan os.Signal, 1)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop
}

// Stop GracefulStop GRPC
func (s *Server) Stop() {
	s.Server.GracefulStop()
	fmt.Println("Server gracefully stopped")
}

// NewServer ...
func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	pingPongCtrl *controller.PingPongController,
	processTracingController *controller_process_tracing.Controller,
// v *validator.CustomValidator,
) *Server {
	option := grpc.ChainUnaryInterceptor(
		// grpc_errors.UnaryServerInterceptor(),
		// validator.UnaryServerInterceptor(v),
	)

	s := &Server{
		Server:                   grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:                   config,
		HealthCtrl:               healthCtrl,
		PingPongCtrl:             pingPongCtrl,
		ProcessTracingController: processTracingController,
	}

	s.Configure()

	return s
}
