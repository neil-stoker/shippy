package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"

	pb "github.com/neil-stoker/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/neil-stoker/shippy/shippy-service-vessel/proto/vessel"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository simulating theuse of a datastore of some
// kind. This will be replaced with a real implementation later.
type Repository struct {
	consignments []*pb.Consignment
}

// Create a new consigment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll consignments from the repository
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures to
// give you a better idea.
type service struct {
	repo repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment creates the consignment
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message created in the protobuf definition
	res.Created = true
	res.Consignment = consignment

	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	vessleResponse, err:= s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n",vesselResponse.Vessel.Name)
	if err!= nil { return err }

	req.VesselId=vessel.Response.Vessel.Id

	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	vesselClient:=vessel.Proto.NewVesselServiceClient("shippy.service.vessel", srv.Client())

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
