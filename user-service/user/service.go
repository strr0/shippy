package user

import (
	"context"
	"go-micro.dev/v4"
	"golang.org/x/crypto/bcrypt"
	pb "user-service/proto/user"
	"user-service/token"
)

type Service struct {
	Repo IRepository
	Publisher micro.Event
}

func (s *Service) Get(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	user, err := s.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (s *Service) GetAll(ctx context.Context, req *pb.UserRequest, res *pb.UserResponse) error {
	users, err := s.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (s *Service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := s.Repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return err
	}
	tk, err := token.Encode(user)
	if err != nil {
		return err
	}
	res.Token = tk
	return nil
}

func (s *Service) Create(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(password)
	if err := s.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	// 推送消息
	if err := s.Publisher.Publish(ctx, req); err != nil {
		return err
	}
	return nil
}

func (s *Service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	decode, err := token.Decode(req.Token)
	if err != nil {
		return err
	}
	res.Token = req.Token
	res.Valid = decode.Valid() == nil
	return nil
}