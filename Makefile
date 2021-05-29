VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)


dev:
	cd pkgs/asynji && air & cd pkgs/pusher && air & cd pkgs/f-manager && air

dev-pusher:
	cd pkgs/pusher && air
	
dev-asynji:
	cd pkgs/asynji && air 

build:
	mkdir build -p
	cd pkgs/asynji && go build -o ../../build/asynji-$(VERSION)
	cd pkgs/pusher && go build -o ../../build/push-service-$(VERSION)
	go test -v ./...

run:
	./build/asynji-$(VERSION) && ./build/push-service-$(VERSION)

docker-build:
	docker-compose build

drun:
	docker-compose up -d