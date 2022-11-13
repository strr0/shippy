package main

import (
	"consignment-service/consignment"
	pb "consignment-service/proto/consignment"
	userProto "consignment-service/proto/user"
	vesselProto "consignment-service/proto/vessel"
	"context"
	"errors"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"strings"

	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//"time"
)

func main() {
	repo := &consignment.Repository{}

	// mongodb连接
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	//if err != nil {
	//	log.Fatalf("mongo connect failed: %v", err)
	//}
	//defer client.Disconnect(ctx)
	//repo := &consignment.MongoRepository{Client: client}

	// 创建委托服务
	srv := micro.NewService(
		micro.Name("shippy.consignment"),
		micro.WrapHandler(AuthWrapper),
	)
	vesselClient := vesselProto.NewVesselService("shippy.vessel", srv.Client())
	srv.Init()
	_ = pb.RegisterShippingServiceHandler(srv.Server(), &consignment.Service{Repo: repo, VesselClient: vesselClient})
	if err := srv.Run(); err != nil {
		log.Fatalln("run failed...")
	}
}

// token验证
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := strings.Replace(meta["Authorization"], "Bearer ", "", 1)
		userService := userProto.NewUserService("shippy.user", client.DefaultClient)
		_, err := userService.ValidateToken(context.Background(), &userProto.Token{Token: token})
		if err != nil {
			return err
		}
		return fn(ctx, req, rsp)
	}
}
