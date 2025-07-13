#!/bin/bash

GO_VERSION=${1:-go1.20}

# Install Go tools
export GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
go env -w GOPROXY=$GOPROXY

go install golang.org/dl/$GO_VERSION@latest
export PATH=$PATH:$HOME/go/bin

$GO_VERSION download

alias go="$GO_VERSION"

echo "Using $(go version)"
echo "Using $(${GO_VERSION} version)"

${GO_VERSION} install golang.org/x/tools/gopls@$GO_VERSION
${GO_VERSION} install github.com/cweill/gotests/...@$GO_VERSION
${GO_VERSION} install github.com/fatih/gomodifytags@$GO_VERSION
${GO_VERSION} install github.com/josharian/impl@$GO_VERSION
${GO_VERSION} install github.com/haya14busa/goplay/cmd/goplay@$GO_VERSION
${GO_VERSION} install github.com/go-delve/delve/cmd/dlv@$GO_VERSION
${GO_VERSION} install github.com/golangci/golangci-lint/cmd/golangci-lint@$GO_VERSION

${GO_VERSION} mod tidy