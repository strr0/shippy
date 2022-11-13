package vessel

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	pb "vessel-service/proto/vessel"
)

const (
	dbName = "shippy"
	vesselCollection = "vessel"
)

type IRepository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
}

// 基于数组
type Repository struct {
	Vessels []*pb.Vessel
}

func (repo *Repository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.Vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

func (repo *Repository) Create(vessel *pb.Vessel) error {
	updated := append(repo.Vessels, vessel)
	repo.Vessels = updated
	return nil
}

// 基于mongo
type MongoRepository struct {
	Client *mongo.Client
}

func (repo *MongoRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	one := repo.Client.Database(dbName).Collection(vesselCollection).FindOne(ctx, bson.M{
		"capacity": bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	})
	res := new(pb.Vessel)
	err := one.Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *MongoRepository) Create(vessel *pb.Vessel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repo.Client.Database(dbName).Collection(vesselCollection).InsertOne(ctx, vessel)
	if err != nil {
		return err
	}
	return nil
}