FILENAME ?= default.proto
generate:
	protoc --go_out=. --go-grpc_out=. $(FILENAME)