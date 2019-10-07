#!/bin/bash

SYSTEMS=(linux windows darwin)
ARCH=amd64

function getCurrentTime() {
    date '+%m-%d-%Y %H:%M:%S'
}

for os in ${SYSTEMS[@]}; do
    echo "[$(getCurrentTime)] build $os binary..."
    GOOS=${os} GOARH=amd64 go build -o check-prom-quey-${os}-${ARCH}
    tar -czf check-prom-query-${os}-${ARCH}.tar.gz check-prom-quey-${os}-${ARCH} && rm -rf check-prom-quey-${os}-${ARCH}
done

echo "[$(getCurrentTime)] Build complete"
