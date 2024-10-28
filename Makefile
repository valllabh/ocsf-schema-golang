.PHONY: build
build:

PROTOC_IMAGE=ocsf-schema-golang-protoc
PROTOC_VERSION=0.1.2

protocontainer:
	docker build --load -t $(PROTOC_IMAGE):$(PROTOC_VERSION) -f dev/protoc/protoc.Dockerfile dev/protoc/

generate-proto: protocontainer
	docker run --rm \
	--mount type=bind,source="$(PWD)",target=/build \
	$(PROTOC_IMAGE):$(PROTOC_VERSION) \
	generate-proto.sh

generate-go: protocontainer
	docker run --rm \
	--mount type=bind,source="$(PWD)",target=/build \
	$(PROTOC_IMAGE):$(PROTOC_VERSION) \
	generate-go.sh

generate-validator:
	go generate ./ocsfvalidate/...

generate: generate-proto generate-go generate-validator
