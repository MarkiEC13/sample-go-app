#!/bin/bash
# watches source update and recompile
set -e

# run db migrations

go run cmd/dbmigrate/main.go

# go run cmd/dbmigrate/main.go -dbname=sampletest

# watches source update and recompile
GO111MODULE=off go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build -o main cmd/api/main.go" --command=./main
