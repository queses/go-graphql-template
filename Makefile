gen:
	go generate ./...

build:
	go build src/server.go

serve: build
	./server

test-docker-up:
	docker-compose -p go-graphql-template -f ./docker/test/docker-compose.yml up -d

test-docker-down:
	docker-compose -p go-graphql-template -f ./docker/test/docker-compose.yml down

migrate:
	go run src/migrate/migrate.go