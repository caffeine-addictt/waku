#!/bin/sh

set -e

rm -rf completions
mkdir completions

for s in bash zsh fish; do
  go run main.go completion "$s" >"completions/waku.$s"
done
