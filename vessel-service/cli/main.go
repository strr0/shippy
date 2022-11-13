package main

import (
	"context"
	"go-micro.dev/v4"
	"log"
	pb "vessel-service/proto/vessel"
)

func main() {
	srv := micro.NewService()
	srv.Init()
	service := pb.NewVesselService("shippy.vessel", srv.Client())
	res, err := service.FindAvailable(context.Background(), &pb.Specification{Capacity: 500, MaxWeight: 200000})
	if err != nil {
		log.Println("find failed...")
		return
	}
	log.Println(res)
}