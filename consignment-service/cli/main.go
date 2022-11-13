package main

import (
	pb "consignment-service/proto/consignment"
	"context"
	"go-micro.dev/v4"
	"go-micro.dev/v4/metadata"
	"log"
)

func main() {
	srv := micro.NewService()
	srv.Init()
	cli := pb.NewShippingService("shippy.consignment", srv.Client())

	// 添加token
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjgzNTAxODksImlkIjoiZTU3NDkxNDUtYTk0My01YzgxLWYzNGItZDMyZWQxZGFiNzU3IiwibmFtZSI6IkpvaG5uaWUifQ.gEqxmnwvde8ff3AZceGFZUbryazv6vHP2oK3BIcBp7A",
	})

	// 创建委托
	//resp, err := cli.CreateConsignment(context.Background(), &pb.Consignment{
	//	Id:          "consignment_id",
	//	Description: "This is a test consignment",
	//	Weight:      55000,
	//	Containers: []*pb.Container{
	//		{CustomerId: "cust001", UserId: "user001", Origin: "Manchester, United Kingdom"},
	//		{CustomerId: "cust002", UserId: "user001", Origin: "Derby, United Kingdom"},
	//		{CustomerId: "cust005", UserId: "user001", Origin: "Sheffield, United Kingdom"},
	//	},
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Printf("create: %v", resp.Created)

	// 获取委托列表
	getAll, err := cli.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(getAll)
}
