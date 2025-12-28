#!/bin/bash
set -e

if command -v swag >/dev/null 2>&1; then
    swag init -g ./cmd/vanish-vault-api/main.go -o api/docs --parseDependency --parseInternal
else
    echo "swag CLI tool is not installed. Please install it by running 'go install github.com/swaggo/swag/cmd/swag@latest'"
    exit 1
fi
