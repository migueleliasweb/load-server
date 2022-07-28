#!/usr/bin/env bash

GOOS=darwin GOARCH=amd64 go build -o load-server-darwin-amd64 && \
GOOS=linux GOARCH=amd64 go build -o load-server-linux-amd64