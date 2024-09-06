#!/bin/bash

# Check deps
HOST_GIT="$(which git)"
if [ -z "$HOST_GIT" ]; then
  echo "git not found"
  exit 1
fi

docker run --rm -it \
  -v "$(which git):/usr/bin/git" \
  -v "$(pwd):/app" \
  waku "$@"
