#!/usr/bin/env bash

PROJECT=foosball-api
BIN=bin
GOOS=linux
GOARCH=amd64

export GO111MODULE=on

for DIR in $(ls cmd); do
    for FILE in $(ls "cmd/${DIR}"); do
        ENDPOINT="${FILE%%.go}"
        echo "Compiling cmd/${DIR}/${FILE} to ${BIN}/${PROJECT}-${DIR}-${ENDPOINT}"
        GOOS="${GOOS}" GOARCH="${GOARCH}" CGO_ENABLED=0 go build -ldflags="-s -w" -o "${BIN}/${PROJECT}-${DIR}-${ENDPOINT}" "cmd/${DIR}/${FILE}"
        chmod 755 "${BIN}/${PROJECT}-${DIR}-${ENDPOINT}"
    done
done