package main

import (
	"github/Jaskerv/go-grpc/pkg/chat"
	"github/Jaskerv/go-grpc/pkg/logger"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	var PORT string

	if PORT = os.Getenv("PORT"); PORT == "" {
		logger.Logger.Fatal("PORT not defined")
	}

	listen, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		logger.Logger.Fatalf("Failed to listen to port %v: %v", PORT, err)
	}

	chatServer := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	logger.Logger.Infof("Starting Server on port: %v 🚀", PORT)
	if err := grpcServer.Serve(listen); err != nil {
		logger.Logger.Fatalf("Failed to serve on port %v: %v", PORT, err)
	}
}
