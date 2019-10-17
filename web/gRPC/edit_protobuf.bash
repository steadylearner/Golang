#!/bin/bash

protoc -I typeDefs/ typeDefs/helloworld.proto --go_out=plugins=grpc:typeDefs

# Use more complicated Bash or Python commands if you have many protobuf files.
