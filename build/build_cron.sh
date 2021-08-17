#!/bin/bash

source /etc/profile

MAIN_PATH=$1

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o crontest "$MAIN_PATH"