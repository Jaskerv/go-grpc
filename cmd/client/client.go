package main

import (
	"context"
	"github/Jaskerv/go-grpc/pkg/chat"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	var PORT string

	if PORT = os.Getenv("PORT"); PORT == "" {
		log.Fatal("PORT not defined")
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}
	defer conn.Close()

	chatService := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := chatService.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}