#!/bin/bash

if [[ $# -ne 2 ]]; then
    echo "usage: ./listen_tcp.sh [PORT] [FILES_NUMBER]"
    exit 92
fi

PORT=$1
NUM_FILES=$2

for i in $(seq 1 $NUM_FILES)
do
   nc -nlvp $PORT > file-$i
done