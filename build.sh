#!/usr/bin/env bash

GOARCH=amd64

for GOOS in linux darwin; do
  echo "Building for $GOOS/$GOARCH.."
  env GOOS=$GOOS GOARCH=$GOARCH go build -o pkgs/shadow_check-${GOOS}_${GOARCH}
done
which upx && upx pkgs/*