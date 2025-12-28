#!/bin/bash
set -e

if command -v sqlc >/dev/null 2>&1; then
    sqlc generate
else
    echo "sqlc CLI tool is not installed. Please install it by running 'go install github.com/kyleconroy/sqlc/cmd/sqlc@latest'"
    exit 1
fi
