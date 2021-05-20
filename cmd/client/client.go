package main

import (
	"context"
	"github/Jaskerv/go-grpc/pkg/chat"
	"github/Jaskerv/go-grpc/pkg/logger"
	"os"

	"google.golang.org/grpc"
)

func main() {
	var SERVER_ADDRESS string

	if SERVER_ADDRESS = os.Getenv("SERVER_ADDRESS"); SERVER_ADDRESS == "" {
		logger.Logger.Fatal("PORT not defined")
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(SERVER_ADDRESS, grpc.WithInsecure())

	if err != nil {
		logger.Logger.Fatalf("Failed to connect: %s", err)
	}
	defer conn.Close()

	chatService := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := chatService.SayHello(context.Background(), &message)
	if err != nil {
		logger.Logger.Fatalf("Error when calling SayHello: %s", err)
	}

	logger.Logger.Infof("Response from Server: %s", response.Body)
}
