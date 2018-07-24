#!/bin/bash

# mac
GOOS=darwin GOARCH=amd64 go build -o build/osccap *.go
# windows
GOOS=windows GOARCH=amd64 go build -o build/osccap.exe *.go

exit 0