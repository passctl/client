#!/bin/bash -e
mkdir -p release

CGO_ENABLED=0 go build
tar czvf release/passctl_linux-amd64.tar.gz client

GOARCH=386 CGO_ENABLED=0 go build
tar czvf release/passctl_linux-386.tar.gz client

rm -f client
