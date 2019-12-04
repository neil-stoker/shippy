package main

import (
	"errors"

	pb "github.com/neil-stoker/shippy/shippy-service-vessel/proto/vessel"
)

// Repository is an interface
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselRepository is a collection of vessels
type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable checks a specification against a list of vessels.
// If capacity and max weight are below a vessels capacity and max weight
// then return that vessel
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}
