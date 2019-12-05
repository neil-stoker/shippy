#!/usr/bin/make

PROTO_NAME=user

%.pb.go: %.protoc
	protoc -I. --go_out=plugins=micro:. proto/$(PROTO_NAME)/$(PROTO_NAME).proto

build: 
	protoc -I. --go_out=plugins=micro:. proto/$(PROTO_NAME)/$(PROTO_NAME).proto



# This is a "phony" target - an alias for the above command, so "make compile"
# still works.
proto: proto/$(PROTO_NAME)/$(PROTO_NAME).pb.go
	protoc -I. go_out=plugins=micro:. proto/$(PROTO_NAME)/$(PROTO_NAME).proto
