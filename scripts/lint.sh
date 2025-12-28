#!/bin/bash
set -e

if command -v golangci-lint >/dev/null 2>&1; then
    golangci-lint run ./...
else
    echo "golangci-lint is not installed. Please install it by following the instructions at https://golangci-lint.run/docs/welcome/install/local/"
    exit 1
fi