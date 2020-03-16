#!/bin/sh
#go mod init
go mod tidy
#go get ./ ...
go mod vendor
go run main.go --help
go run main.go --province roma
go run main.go --region lazio
