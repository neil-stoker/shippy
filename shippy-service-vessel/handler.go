package main

import (
	"context"

	pb "github.com/neil-stoker/shippy/shippy-service-vessel/proto/vessel"
	"gopkg.in/mgo.v2"
)

type handler struct {
	session *mgo.Session
}

// GetRepo -
func (s *handler) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

// Create -
func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.GetRepo().Create(req); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true

	return nil
}

func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// Find the next available vessel
	vessel, err := s.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as pat of the response message type
	res.Vessel = vessel
	return nil
}
