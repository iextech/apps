#!/bin/bash

mkdir -p pkg/ansibleapp
protoc --go_out=. --go-grpc_out=. proto/ansibleapp.proto

mv github.com/iextech/apps/pkg/ansibleapp/*.go pkg/ansibleapp/
rm -rf github.com
