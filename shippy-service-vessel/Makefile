#!/usr/bin/make
PB = $(wildcard proto/vessel/*.proto)
GO = $(PB:.proto=.pb.ho)

%.pb.go: %.proto
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto

build: proto
	docker build -t shippy-service-vessel .

run:
	docker run -p 50052: 50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-vessel

proto: proto/vessel/vessel.pb.go
