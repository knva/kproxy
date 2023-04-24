@echo off
set GOOS=linux
set GOARCH=amd64
go build -o kproxy main.go
echo Compilation complete: output-file
