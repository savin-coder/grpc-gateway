#!/bin/sh
# Demonstrate how read actually works
echo generating all protos
   protoc -I ./proto \
      --go_out ./generated --go_opt paths=source_relative \
      --go-grpc_out ./generated --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./generated --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative\
      ./proto/ecard/*proto


