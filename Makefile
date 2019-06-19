export PATH := $(shell go env GOPATH)/bin:dep/protoc/bin:$(PATH)

PROTO_TARGETS = pkg/greet/greet.pb.go

# --------------------------

dep/protoc/bin/protoc:
	mkdir -p dep/protoc
	curl -L -o dep/protoc/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip
	cd dep/protoc; unzip -o protoc.zip

dep-install: dep/protoc/bin/protoc
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get ./...

# --------------------------

$(PROTO_TARGETS): %.pb.go: %.proto
	protoc -I $(dir $@) $< --go_out=plugins=grpc:$(dir $@)

proto: $(PROTO_TARGETS)

# --------------------------

test: proto
	go test -v ./...

# --------------------------

greetserver: proto
	CGO_ENABLED=0 go build -o output/greetserver ./cmd/greetserver

greetclient: proto
	CGO_ENABLED=0 go build -o output/greetclient ./cmd/greetclient

# --------------------------

.PHONY: dep-install proto test oneserver client
