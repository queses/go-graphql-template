generate:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen
	go get github.com/99designs/gqlgen

build:
	go build src/server.go

serve: build
	./server