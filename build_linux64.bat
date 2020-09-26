
@echo off

set GOOS=linux
set GOARCH=amd64

echo "Build For ws-server ..."

go build -mod=vendor -ldflags "-s -w" -o ws-server

echo "--------- Build For ws-server Success!"

