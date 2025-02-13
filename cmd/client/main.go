package main

import (
	"context"
	"github.com/dankru/GRPC/proto/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.NewClient(":9000", opts...)
	if err != nil {
		log.Fatalf("Не удалось установить соединение: %s", err.Error())
	}

	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{Message: "Ниндзяго"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Status:", response.Status)
}
