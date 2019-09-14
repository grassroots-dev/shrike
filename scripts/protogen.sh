#!/bin/bash
rm -Rf ./service
mkdir service
protoc --proto_path=. --proto_path=third_party --go_out=plugins=grpc:service ./shrike.proto