ARG BUILD_BASE_IMAGE="golang:1.25-alpine"
ARG FINAL_BASE_IMAGE="alpine:latest"

FROM --platform=$BUILDPLATFORM ${BUILD_BASE_IMAGE} AS base
WORKDIR /app
RUN apk add --no-cache git make ca-certificates
COPY go.mod go.sum ./
RUN go mod download

FROM base AS dev
RUN go install github.com/air-verse/air@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.7.2
COPY . .
CMD ["air", "-c", ".air.toml"]

FROM base AS builder
ARG TARGETOS
ARG TARGETARCH
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o vanish-vault-api ./cmd/vanish-vault-api

FROM ${FINAL_BASE_IMAGE} AS prod
WORKDIR /app
COPY --from=builder /app/vanish-vault-api .
ENTRYPOINT ["./vanish-vault-api"]