#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -o pkgs/shadow_check-linux_amd64
env GOOS=darwin GOARCH=amd64 go build -o pkgs/shadow_check-darwin_amd64