#!/bin/bash
if [ ! -e /usr/local/bin/protoc ]; then
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.10.0-rc1/$(echo $PROTOC_ZIP)
unzip -o $(echo $PROTOC_ZIP) -d /usr/local bin/protoc
rm -f $(echo $PROTOC_ZIP)
fi