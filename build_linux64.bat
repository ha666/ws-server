
@echo off

set GOOS=windows
set GOARCH=amd64

echo "Build For ws-server ..."

go build -mod=vendor -ldflags "-s -w" -o ws-server.exe

echo "--------- Build For ws-server Success!"

