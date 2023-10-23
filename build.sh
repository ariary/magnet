#!/usr/bin/env bash

if [[ $# -ne 4 ]]; then
    echo "usage: ./build.sh \$FILES \$ENDPOINT \$KEY \$METHOD"
    exit 92
fi

# TARGET_OS=$1
# FILES=$2
# ENDPOINT=$3
# KEY=$4
# METHOD=$5

EXENAME="magnet"
TARGET=$(go tool dist list|gum filter --placeholder="choose target os & arch")
export GOOS=$(echo $TARGET|cut -f1 -d '/')
export GOARCH=$(echo $TARGET|cut -f2 -d '/')

export KEY=$3
export FILES=$(cat $1 | lobfuscator $KEY)
export ENDPOINT=$(echo "$2" | lobfuscator $KEY)
export METHOD=$4

echo "build ${EXENAME}-${GOOS}-${GOARCH} in ${PWD}"
CGO_ENABLED=0 go build -ldflags "-X 'main.FileList=$FILES' -X 'main.Key=$KEY' -X 'main.Endpoint=$ENDPOINT' -X 'main.Method=$METHOD'" -o ${EXENAME}-${GOOS}-${GOARCH} cmd/magnet/magnet.go
