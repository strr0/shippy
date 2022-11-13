package main

import (
	"context"
	pb "email-service/proto/user"
	"go-micro.dev/v4"
	"log"
)

var topic = "user.created"

type Subscriber struct {}

func (Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Printf("Sending email to: %s", user.Name)
	return nil
}

func main() {
	srv := micro.NewService()
	srv.Init()
	_ = micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))
	if err := srv.Run(); err != nil {
		log.Fatalln("run failed...")
	}
}