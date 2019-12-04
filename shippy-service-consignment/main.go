package main

import (
	"context"
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/neil-stoker/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/neil-stoker/shippy/shippy-service-vessel/proto/vessel"
)

const (
	port        = ":50051"
	defaultHost = "mongodb://datastore:27017"
)

func main() {

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
		// micro.Version("latest"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.client", srv.Client())
	h := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
