package consignment

import (
	pb "consignment-service/proto/consignment"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignment"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// 基于切片
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// 基于mongo
type MongoRepository struct {
	Client *mongo.Client
}

func (repo *MongoRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repo.Client.Database(dbName).Collection(consignmentCollection).InsertOne(ctx, consignment)
	if err != nil {
		return nil, err
	}
	return consignment, nil
}

func (repo *MongoRepository) GetAll() []*pb.Consignment {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	find, err := repo.Client.Database(dbName).Collection(consignmentCollection).Find(ctx, bson.D{})
	if err != nil {
		return nil
	}
	defer find.Close(ctx)
	var res []*pb.Consignment
	err = find.All(ctx, &res)
	if err != nil {
		return nil
	}
	return res
}
