dev:
	cd pkgs/asynji && air & cd pkgs/pusher && air

build:
	mkdir build
	cd pkgs/asynji && go build -o ../../build/asynji
	cd pkgs/pusher && go build -o ../../build/push-service 
	go test -v ./...

run:
	./build/asynji && ./build/push-service

docker-build:
	docker-compose build

drun:
	docker-compose up -d