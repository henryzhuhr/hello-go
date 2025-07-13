#!/bin/bash

# 编译 dockerfiles/Dockerfile

docker build \
  -t hello-go:latest \
  -f dockerfiles/Dockerfile .

# Optional build args
# --build-arg GO_VERSION=1.23.0 \
# --build-arg BASE_INSTALL_LIST=ubuntu-server \
# --build-arg EXTENT_INSTALL_LIST="vim nano tree" \
# --build-arg MIRRORS_URL=mirrors.cloud.tencent.com \
# --build-arg CLEAN_APT_CACHE=1 \