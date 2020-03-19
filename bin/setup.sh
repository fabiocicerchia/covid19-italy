#!/bin/bash
git submodule init
git submodule update

# CLI
go mod tidy
go mod vendor
go build main.go +o ./bin/covid19-trend
./bin/covid19-trend --help

# WEB
echo "Show data http://127.0.0.1:8001/index.html"
php -S 127.0.0.1:8081 -t public
