#!/bin/bash

echo -n "Enter key: " 
read key
echo -n "Enter filename containing files to exfiltrate: "
read files
echo -n "Enter endpoint: "
read endpoint
echo -n "Enter method (http, tcp): "
read method

export KEY="$key"
export FILES="$files"
export ENDPOINT="$endpoint"
export METHOD="$method"