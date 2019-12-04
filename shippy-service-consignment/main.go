package main

import (
	"context"
	"fmt"
	"log"

	micro "github.com/micro/go-micro"

	pb "github.com/neil-stoker/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/neil-stoker/shippy/shippy-service-vessel/proto/vessel"
)

const (
	port = ":50051"
)

// Repository holds the consignment data
type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentRepository - Dummy repository simulating theuse of a datastore of some
// kind. This will be replaced with a real implementation later.
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

// Create a new consigment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll consignments from the repository
func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures to
// give you a better idea.
type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment creates the consignment
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// Set the VesselId as the vessel we got back from the vessel service
	req.VesselId = vesselResponse.Vessel.Id

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
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	repo := &ConsignmentRepository{}

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
		micro.Version("latest"),
	)

	srv.Init()

	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
