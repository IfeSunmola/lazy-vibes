#!/usr/bin/env bash

# linux
env GOOS=linux GOARCH=amd64 go build -o builds/tracker-linux-amd64 tracker.go
env GOOS=linux GOARCH=arm64 go build -o builds/tracker-linux-arm64 tracker.go

# windows
env GOOS=windows GOARCH=amd64 go build -o builds/tracker-win-amd64.exe tracker.go
env GOOS=windows GOARCH=arm64 go build -o builds/tracker-win-arm64.exe tracker.go


env GOOS=darwin GOARCH=amd64 go build -o builds/tracker-mac-amd64 tracker.go
env GOOS=darwin GOARCH=arm64 go build -o builds/tracker-mac-arm64 tracker.go
