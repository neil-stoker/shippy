#!/usr/bin/make

containers=`docker ps | grep 'shippy' | sed 's/ / /' | awk '{print $1}' |  tr  '\n' ' '`

build:
	@docker-compose build | grep "Successfully tagged" --color=always

start:
	docker-compose up -d shippy-service-vessel
	docker-compose up -d shippy-service-consignment

stop:
	# echo "containers .$(containers)."
	# `docker ps| grep shippy | awk '{print $1}'` | docker stop
	docker stop  $(docker ps -qa)

cli:
	docker-compose run shippy-cli-consignment


