
#!/bin/bash

build_dir=`pwd`

echo $(date +"%H:%M:%S")  "start compile"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags "-s -w" -o ws-server
echo $(date +"%H:%M:%S")  "finish compile"
