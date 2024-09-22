ARG GO_VERSION=1.23

FROM mcr.microsoft.com/devcontainers/go:1-${GO_VERSION}-bookworm

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod \
    go install github.com/spf13/cobra-cli@v1.3.0 \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 \
    && go install github.com/goreleaser/goreleaser/v2@v2.2.0 \
    && go install golang.org/x/vuln/cmd/govulncheck@v1.1.3
