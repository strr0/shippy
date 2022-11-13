package vessel

import (
	"context"
	pb "vessel-service/proto/vessel"
)

type Service struct {
	Repo IRepository
}

func (s *Service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.VesResponse) error {
	vessel, err := s.Repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func (s *Service) Create(ctx context.Context, req *pb.Vessel, res *pb.VesResponse) error {
	err := s.Repo.Create(req)
	if err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}