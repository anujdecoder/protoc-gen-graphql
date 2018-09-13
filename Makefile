
LOCAL_PATH = $(shell pwd)

.PHONY: test
example: install build-proto-example
	protoc -I /usr/local/include \
    -I ${LOCAL_PATH} \
    -I ${GOPATH}/src \
    --graphql_out=plugins:. \
    example/**.proto

.PHONY: install
install: packr
	go install .

.PHONY: packr
packr:
	cd templates; packr

.PHONY: build-proto-example
build-proto-example:
	protoc -I /usr/local/include \
        -I ${LOCAL_PATH} \
        -I ${GOPATH}/src \
        --go_out=plugins=grpc:. \
        example/**.proto

build-graphql-proto:
	protoc -I /usr/local/include \
		-I ${LOCAL_PATH} \
		--go_out=paths=source_relative,plugins=:. \
		graphql/*.proto