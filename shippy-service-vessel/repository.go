package main

import (
	pb "github.com/neil-stoker/shippy/shippy-service-vessel/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

// Repository is an interface
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

// VesselRepository is a collection of vessels
type VesselRepository struct {
	session *mgo.Session
}

// Close -
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

// collection -
func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}

// Create -
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// FindAvailable checks a specification against a list of vessels.
// If capacity and max weight are below a vessels capacity and max weight
// then return that vessel
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}

	return vessel, nil
}
