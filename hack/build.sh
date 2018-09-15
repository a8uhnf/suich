#!/usr/bin/env bash

$BINARY_NAME="suich"
$TAG="1.0"

echo $BINARY_NAME
rm $BINARY_NAME
goos=$()
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o suich ./server/main.go
docker build -t a8uhnf/grpc-server:1.0 .
rm $BINARY_NAME
docker push a8uhnf/$BINARY_NAME:$TAG