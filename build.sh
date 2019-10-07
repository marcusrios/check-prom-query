#!/bin/bash

SYSTEMS=(linux windows darwin)
ARCH=amd64

function getCurrentTime() {
    date '+%m-%d-%Y %H:%M:%S'
}

for os in ${SYSTEMS[@]}; do
    echo "[$(getCurrentTime)] build $os binary..."
    GOOS=${os} GOARH=amd64 go build -o check-prom-query-${os}-${ARCH}
    tar -czf check-prom-query-${os}-${ARCH}.tar.gz check-prom-query-${os}-${ARCH} && rm -rf check-prom-query-${os}-${ARCH}
done

echo "[$(getCurrentTime)] Build complete"
