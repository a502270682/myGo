#!/bin/bash

source /etc/profile

VERSION=$(git describe --tags 2>/dev/null)
COMMIT=$(git rev-parse --short HEAD)
TIME=$(date +%FT%T)
MAIN_PATH=$1
if [ -z $VERSION ]; then
    VERSION=$COMMIT
fi

VerPkg="myGo"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o myGo -ldflags "-X ${VerPkg}.Version=$VERSION -X ${VerPkg}.GitCommit=$COMMIT -X ${VerPkg}.BuildTime=${TIME}" "$MAIN_PATH"
