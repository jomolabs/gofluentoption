#!/bin/bash

[ ! -d "bin/" ] && mkdir bin/

function getVersion() {
    local version="$(git rev-list -1 HEAD 2>/dev/null)"
    if [ $? -ne 0 -o -z "$version" ]; then
        version="local.$(hostname).$(date "+%Y-%m-%d@%H:%M:%S")"
    fi

    echo "$version"
}

version="$(getVersion)"
>&2 echo "Building version \"$version\"..."
go mod vendor
go build -ldflags "-X main.version=$version" -o bin/gofluentoption gofluentoption.go
