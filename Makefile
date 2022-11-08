PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

generate:
	protoc --go_out=. --go-grpc_out=. --go_opt=module=${PACKAGE} --go-grpc_opt=module=${PACKAGE} proto/*.proto

build:
	go build -o bin/server server/main.go
	go build -o bin/client client/main.go client/send_source.go client/send_source_stream.go
	docker build --tag docker-grpc .

test:
	go test ./tests

deploy:
	docker run -d -p 8001:50051 docker-grpc
