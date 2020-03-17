#!/bin/sh
go test -v -coverprofile=c.out test/valueobject/userchoice/userchoice_test.go
go tool cover -html=c.out

go test -v -coverprofile=c.out test/utils/getoptshandler/getoptshandler_test.go
go tool cover -html=c.out
