package main

import (
	"context"

	pb "github.com/neil-stoker/shippy/shippy-service-consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository -
type Repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (mr *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := mr.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll -
func (mr *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment

	cur, err := mr.collection.Find(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err = cur.Decode(&consignment); err != nil {
			return nil, err
		}

		consignments = append(consignments, consignment)
	}

	return consignments, nil
}
