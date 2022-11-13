package main

import (
	//"context"
	"go-micro.dev/v4"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//"time"
	pb "vessel-service/proto/vessel"
	"vessel-service/vessel"
)

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	repo := &vessel.Repository{Vessels: vessels}

	// mongodb连接
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	//if err != nil {
	//	log.Fatalf("mongo connect failed: %v", err)
	//}
	//defer client.Disconnect(ctx)
	//repo := &vessel.MongoRepository{Client: client}

	// 代销服务
	srv := micro.NewService(
		micro.Name("shippy.vessel"),
	)
	srv.Init()
	_ = pb.RegisterVesselServiceHandler(srv.Server(), &vessel.Service{Repo: repo})
	if err := srv.Run(); err != nil {
		log.Fatalln("run failed...")
	}
}