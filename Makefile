gen:
	go generate ./...

build:
	go build src/server.go

serve: build
	./server
