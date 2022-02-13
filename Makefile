GEN_DOCS_PATH := ./docs
GEN_DOCS_NAME := kv-svc.swagger.json
GEN_GO_PATH := ./pkg/gen
GEN_GW_PATH := ${GEN_GO_PATH}
PACKAGE_PATH := /github.com$(shell pwd | sed -n -e 's/^.*github.com//p' | sed -e 's/-/_/g')

help:
	echo Create and clean up generated files in addition to building & testing application

clean:
	rm -rf ${GEN_GO_PATH} ${GEN_DOCS_PATH}

install:
	# Make sure GOBIN is set correctly:
	# go env -w GOBIN=$$GOPATH/bin
	# rm -rf vendor/
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go mod vendor -v

generate:
	mkdir -p ${GEN_GO_PATH} ${GEN_GW_PATH} ${GEN_DOCS_PATH}
	protoc -I/usr/local/include -I. \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2 \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=${GEN_GO_PATH} \
		--go_opt paths=source_relative \
		--go-grpc_out=${GEN_GW_PATH} \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out=logtostderr=true,allow_delete_body=true:${GEN_GW_PATH} \
		--openapiv2_out=logtostderr=true,allow_delete_body=true:${GEN_DOCS_PATH} \
		./*.proto
	mv ${GEN_GW_PATH}${PACKAGE_PATH}/* ${GEN_GW_PATH}
	rm -rf ${GEN_GW_PATH}/github.com

build:
	go build -o kv-store-grpc

rpc: build
	./kv-store-grpc rpc

rest: build
	./kv-store-grpc rest

test:
	curl -X POST localhost:8080/v1/kv -d '{"key": "nice", "value": "gotem" }' && echo
	curl -X GET localhost:8080/v1/kv/nice && echo
	curl -X DELETE localhost:8080/v1/kv/nice && echo
	# I should expect an error here...
	curl -X GET localhost:8080/v1/kv/nice && echo
	# I should expect an error here...
	curl -X DELETE localhost:8080/v1/kv/nice && echo

PHONY: help clean generate rest rpc build test install
