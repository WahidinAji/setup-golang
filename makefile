# run bash commands to compile and run the program

dev:
	sh ./bashscript.sh

test:
	go test -v ./...

build:
	go build -o bin/ ./...


# Docker commands
docker-dev:
	docker-compose up -d
	
docker-stop:
	docker-compose down

docker-build:
	docker-compose build

docker-restart:
	docker-compose down
	docker-compose up -d

docker-logs:
	docker-compose logs -f
