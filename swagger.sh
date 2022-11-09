#!/bin/bash
# Demonstrate how read actually works
echo generating all swagger for protos
     protoc -I ./proto \
         --openapiv2_out ./openapiv2 \
         --openapiv2_opt logtostderr=true \
         ./proto/ecard/*proto

