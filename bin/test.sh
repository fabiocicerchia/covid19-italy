#!/bin/sh
go test -coverpkg=./... -coverprofile cover.out ./...
go tool cover -html=cover.out -o coverage.html
rm cover.out