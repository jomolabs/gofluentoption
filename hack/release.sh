#!/bin/bash

mkdir -p release/bin/
rm -f release/bin/* release/*.tar.gz

onPlatformName="gofluentoption.$(uname -s | tr '[[:upper:]]' '[[:lower:]]')-$(uname -m)"
OUTPUT_NAME=${onPlatformName} OUTPUT_PATH=release/bin/ ./hack/build.sh
tar cvzf release/$onPlatformName.tar.gz -C release/bin $onPlatformName

for buildImage in 'debian:stretch' alpine; do
    imageOption=
    osOption=

    if [ -n "$(echo $buildImage | grep ':')" ]; then
        osOption="$(echo $buildImage | cut -d ':' -f1)"
        imageOption="$(echo $buildImage | cut -d ':' -f2)"
    else
        osOption=$buildImage
        imageOption=$buildImage
    fi

    docker run \
        --rm \
        -v $(pwd):/go/src/github.com/jomolabs/gofluentoption \
        -v $(pwd)/release/bin:/output \
        --workdir /go/src/github.com/jomolabs/gofluentoption \
        golang:1.13.6-$imageOption \
        go build -o /output/gofluentoption.$osOption-x86_64 gofluentoption.go

    tar cvzf release/gofluentoption.${osOption}-x86_64.tar.gz -C release/bin gofluentoption.${osOption}-x86_64
done
