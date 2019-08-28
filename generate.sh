#!/bin/bash

protoc grpc.proto --go_out=plugins=grpc:.
