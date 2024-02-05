#!/usr/bin/env bash

go install golang.org/x/tools/cmd/goimports@latest

# install golangci-lint
if ! command -v golangci-lint &> /dev/null; then
  echo "Installing golangci-lint"
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
fi
