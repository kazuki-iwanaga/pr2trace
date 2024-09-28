ARG GO_VERSION=1.23

FROM mcr.microsoft.com/devcontainers/go:1-${GO_VERSION}-bookworm

USER vscode

RUN go install github.com/evilmartians/lefthook@v1.7.15
RUN go install golang.org/x/vuln/cmd/govulncheck@v1.1.3
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

RUN go install go.opentelemetry.io/collector/cmd/builder@v0.110.0 \
    && mv /go/bin/builder /go/bin/ocb
