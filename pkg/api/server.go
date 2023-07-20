package api

import (
	"log"
	"net"

	"github.com/athunlal/bookNow-Account-Services/pkg/api/handler"
	"github.com/athunlal/bookNow-Account-Services/pkg/pb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	App *fiber.App
}

func NewServerHttp(userHandler *handler.UserHandler) *ServerHttp {
	app := fiber.New()

	go NewGRPCServer(userHandler, "8890")

	return &ServerHttp{
		App: app,
	}
}

func NewGRPCServer(userHandler *handler.UserHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen to GRPC Port", err)
	}

	// Creating a new GRPC Server
	grpcServer := grpc.NewServer()

	pb.RegisterProfileManagementServer(grpcServer, userHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Could not serve the grpc Server", err)
	}
}

func (s *ServerHttp) Start() {
	err := s.App.Listen(":7778")
	if err != nil {
		log.Fatalf("Error starting Fiber server: %v", err)
	}
}
