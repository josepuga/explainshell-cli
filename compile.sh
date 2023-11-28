#!/bin/bash
# By José Puga. 2023.
app_name=expl
build_options=-ldflags="-s -w"
go fmt && go mod tidy && \
go build "$build_options" -o bin/$app_name . && \
GOOS=windows GOARCH=386 go build "$build_options" -o bin/${app_name}32.exe .
GOOS=windows GOARCH=amd64 go build "$build_options" -o bin/$app_name.exe .