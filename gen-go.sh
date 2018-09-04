#!/bin/bash
set -eo pipefail
set -x

# Path to proto files
PROTO_PATH=.
PROTO_FILE=$PROTO_PATH/*
DOCKER_IMAGE=consensus-protoc

# Clean up previously generated
[ -e "$PROTO_FILE.pb.*" ] && rm $PROTO_FILE.pb.*

docker build -t $DOCKER_IMAGE -f proto.Dockerfile .

# Generates grpc server code
docker run -v $PWD:/tmp -w /tmp $DOCKER_IMAGE protoc \
    --lint_out=. \
    --go_out=plugins=grpc,paths=source_relative:. \
    $PROTO_FILE.proto