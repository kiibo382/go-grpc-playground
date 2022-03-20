.PHONY: install-protobuf
install-protobuf:
	brew install protobuf

.PHONY: setup
setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

.PHONY: gen-proto
gen-proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./pb/rock-paper-scissors.proto

.PHONY: up-server
up-server:
	cd cmd/api && go run main.go

.PHONY: up-client
up-client:
	cd cmd/cli && go run main.go

