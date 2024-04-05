#!/bin/bash
set -euo pipefail

export GOOS=linux

for BIN in detect build; do
  echo Building $BIN
  go build $@ -ldflags="-s -w" -o ./bin/$BIN ./cmd/$BIN/main.go
done
