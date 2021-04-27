dev:
	cd src && air & cd lib/push-service && air

build:
	mkdir build
	cd src && go build -o ../build/asynji
	cd lib/push-service && go build -o ../../build/push-service 
	go test -v ./...

run:
	./build/asynji && ./build/push-service

docker-build:
	docker-compose build

drun:
	docker-compose up -d