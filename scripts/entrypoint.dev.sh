#!/bin/bash
# watches source update and recompile
set -e

GO111MODULE=off go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build -o main cmd/api/main.go" --command=./main
