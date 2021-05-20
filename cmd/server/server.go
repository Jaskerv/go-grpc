package main

import (
	"github/Jaskerv/go-grpc/pkg/chat"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	var PORT string

	if PORT = os.Getenv("PORT"); PORT == "" {
		log.Fatal("PORT not defined")
	}

	listen, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("Failed to listen to port %v: %v", PORT, err)
	}

	chatServer := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	log.Printf("Starting Server on port: %v 🚀", PORT)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve on port %v: %v", PORT, err)
	}
}
