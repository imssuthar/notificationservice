package main

import (
	"context"
	"imssuthar/notificationsystem/dao"
	"imssuthar/notificationsystem/rest"
	"log"
)

func main() {
	if err := dao.Connect(context.Background()); err != nil {
		log.Fatalf("mongo: %v", err)
	}
	defer func() {
		if err := dao.Disconnect(context.Background()); err != nil {
			log.Printf("mongo disconnect: %v", err)
		}
	}()

	if err := rest.Attach(); err != nil {
		log.Printf("server stopped: %v", err)
	}
}
