FROM golang:1.22-bookworm as build

FROM build as build-go
# https://github.com/protocolbuffers/protobuf-go/releases
ARG PROTOC_GEN_GO_VERSION=1.36.7
ARG PROTOC_GEN_GO_HASH="f3be1721420f0524ed036e16b5b53f13d10052741a7061db9b13f0a4a469d817"


RUN mkdir /xsrc && cd /xsrc && \
    curl -fLO https://github.com/protocolbuffers/protobuf-go/archive/refs/tags/v${PROTOC_GEN_GO_VERSION}.tar.gz && \
    echo "${PROTOC_GEN_GO_HASH} v${PROTOC_GEN_GO_VERSION}.tar.gz" | sha256sum -c - && \
    tar -xzf v${PROTOC_GEN_GO_VERSION}.tar.gz && \
    cd protobuf-go-${PROTOC_GEN_GO_VERSION}/cmd/protoc-gen-go && \
    go install -v . && \
    cd / && \
    rm -rf "/xsrc"

FROM build as build-ocsf-tool

ARG OCSF_TOOL_REPO=https://github.com/pquerna/ocsf-tool
ARG OCSF_TOOL_COMMIT=f8ed1da78a8c36def4584cc0fc7819ab9fd276f6

RUN mkdir /xsrc && cd /xsrc && \
    git clone ${OCSF_TOOL_REPO} && \
    cd ocsf-tool && \
    git checkout ${OCSF_TOOL_COMMIT} && \
    go install -v . && \
    cd / && \
    rm -rf "/xsrc"

# https://github.com/bufbuild/buf/releases
FROM bufbuild/buf:1.56.0 as buf

FROM debian:bookworm as final

RUN mkdir /build && apt-get -y update && apt-get -y install \
    protobuf-compiler \
    jq \
    ca-certificates \
    rsync

COPY --from=buf /usr/local/bin/buf /usr/local/bin
COPY --from=build-go /go/bin/protoc-gen-go /usr/local/bin
COPY --from=build-ocsf-tool /go/bin/ocsf-tool /usr/local/bin
COPY ./generate-proto.sh /usr/local/bin/generate-proto.sh
COPY ./generate-go.sh /usr/local/bin/generate-go.sh
COPY ./config-v*.yaml /configs/

WORKDIR /build
