#!/bin/bash
set -euo pipefail

function no-main() {
  # Check if the git branch is main
  local current_branch=$(git rev-parse --abbrev-ref HEAD)
  if [ "$current_branch" == "main" ]; then
    echo "Can't push main" > 2
    exit 1
  fi
}

PUSH=
set +u
if [ "$1" == "push" ]; then
    PUSH=yes
    shift
fi
set -u

export GOOS=linux

for BIN in detect build; do
  echo Building $BIN
  go build $@ -ldflags="-s -w" -o ./bin/$BIN ./cmd/$BIN/main.go
done

tar czf package/juvo-poetry-buildpack.tar.gz bin buildpack.toml

if [ "$PUSH" == "yes" ]; then
  no-main
  set -x
  git add package/juvo-poetry-buildpack.tar.gz
  git commit --amend --no-edit
  git push --force
fi
