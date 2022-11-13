package main

import (
	"context"
	"go-micro.dev/v4"
	"log"
	pb "user-service/proto/user"
)

func main() {
	srv := micro.NewService()
	srv.Init()
	cli := pb.NewUserService("shippy.user", srv.Client())

	//创建用户
	_, err := cli.Create(context.Background(), &pb.User{
		Name: "Johnnie",
		Email: "jn666@example.com",
		Password: "pass123",
		Company: "go-micro",
	})
	if err != nil {
		log.Println("create failed...")
		return
	}
	log.Println("create success...")

	// 获取用户列表
	//all, err := cli.GetAll(context.Background(), &pb.UserRequest{})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(all)

	// 生成token
	//auth, err := cli.Auth(context.Background(), &pb.User{
	//	Email: "jn666@example.com",
	//	Password: "pass123",
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(auth)
}