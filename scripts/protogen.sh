#!/bin/bash
rm  ./service/shrike.pb.go
protoc --proto_path=. --proto_path=third_party --go_out=plugins=grpc:service ./shrike.proto