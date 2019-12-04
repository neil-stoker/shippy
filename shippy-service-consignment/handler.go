package main

import (
	"context"
	"log"

	pb "github.com/neil-stoker/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/neil-stoker/shippy/shippy-service-vessel/proto/vessel"
)

type handler struct {
	Repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - takes a context and request and creates a consignment
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	log.Printf("Found vessel: %s\n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err = s.Repository.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments -
func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.Repository.GetAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
