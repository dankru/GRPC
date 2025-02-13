package main

import (
	"context"
	"fmt"
	"github.com/dankru/GRPC/proto/notification"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	// Обязательный метод для будущей совместимости grpc, не несёт функционала
	notification.UnimplementedNotificationServiceServer
}

func (s *server) Notify(ctx context.Context, n *notification.NotificationRequest) (*notification.NotificationResponse, error) {
	fmt.Println("RECEIVED NOTIFICATION:", n.Message)
	return &notification.NotificationResponse{Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	notification.RegisterNotificationServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
