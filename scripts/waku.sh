#!/bin/bash

if ! command -v docker &>/dev/null; then
  echo "Error: Docker is not installed. Please install Docker and try again."
  exit 1
fi

IMG="caffeinec/waku:latest"

# Check if GHCR is preferred over DockerHub
if [[ "$1" == "ghcr" ]]; then
  IMG="ghcr.io/caffeine-addictt/waku:latest"
  shift # Remove 'ghcr' from the args
fi

docker run -v "$(pwd):/app" "$IMG" "$@"
