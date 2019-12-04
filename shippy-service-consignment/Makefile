
PB = $(wildcard proto/consigment/*.proto)
GO = $(PB:.proto=.pb.go)

%.pb.go: %.proto
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto

build: proto
	docker build -t shippy-service-consignment .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		shippy-service-consignment
