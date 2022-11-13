package consignment

import (
	pb "consignment-service/proto/consignment"
	vesselProto "consignment-service/proto/vessel"
	"context"
	"log"
)

type Service struct {
	Repo         IRepository
	VesselClient vesselProto.VesselService
}

func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	vesselResponse, err := s.VesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err == nil && vesselResponse != nil && vesselResponse.Vessel != nil {
		log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	}
	if err != nil {
		return err
	}
	req.VesselId = vesselResponse.Vessel.Id
	consignment, err := s.Repo.Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = consignment
	return nil
}

func (s *Service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments := s.Repo.GetAll()
	resp.Consignments = consignments
	return nil
}
