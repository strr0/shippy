package main

import (
	"go-micro.dev/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	pb "user-service/proto/user"
	"user-service/user"
)

var topic = "user.created"

func main() {
	db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/shippy?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Println("gorm failed...")
		return
	}
	//db.AutoMigrate(&pb.User{})
	repo := &user.Repository{Db: db}
	srv := micro.NewService(
		micro.Name("shippy.user"),
	)
	srv.Init()
	// 注册事件驱动
	publisher := micro.NewEvent(topic, srv.Client())
	// 注册用户服务
	_ = pb.RegisterUserServiceHandler(srv.Server(), &user.Service{Repo: repo, Publisher: publisher})
	if err := srv.Run(); err != nil {
		log.Fatalln("run failed...")
	}
}