#!/bin/sh
go mod tidy
go mod vendor
go run main.go --help
go run main.go --province roma
go run main.go --region lazio
./bin/test.sh
