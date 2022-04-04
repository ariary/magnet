#!/bin/bash

if [[ $# -ne 4 ]]; then
    echo "usage: ./build.sh \$TARGET_OS \$FILES \$ENDPOINT \$KEY"
    exit 92
fi

# TARGET_OS=$1
# FILES=$2
# ENDPOINT=$3
# KEY=$4

export TARGET_OS=$1
export KEY=$4
export FILES=$(cat $2 | lobfuscator $KEY)
export ENDPOINT=$(echo "$3" | lobfuscator $KEY)

GOOS=$TARGET_OS GOARCH=amd64 go build -ldflags "-X 'main.FileList=$FILES' -X 'main.Key=$KEY' -X 'main.Endpoint=$ENDPOINT'" cmd/magnet/magnet.go