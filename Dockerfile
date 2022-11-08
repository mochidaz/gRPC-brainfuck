FROM golang:1.18 AS base

# define timezone
ENV TZ Asia/Jakarta

# define work directory
WORKDIR /app

# copy the sourcecode
COPY . .

# install protoc
RUN apt-get update && apt-get install -y protobuf-compiler

# install all go dependencies
RUN cd ~ && export GO111MODULE=auto && go install github.com/envoyproxy/protoc-gen-validate@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest

# generate protocol buffers
RUN make generate

# build beedoor exec
RUN cd /app && go mod vendor && make build

# EXPOSE 8080 is the port that the REST API will be exposed on
EXPOSE 50051
# EXPOSE 8080 is the port that the GRPC will be exposed on. But if deployed in cloud run just use the 8080 port

CMD [ "./bin/server" ]