ARG GO_VERSION=1.23

FROM golang:${GO_VERSION} AS devcontainer

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod \
    go install github.com/spf13/cobra-cli@v1.3.0 \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 \
    && go install github.com/goreleaser/goreleaser/v2@v2.2.0 \
    && go install golang.org/x/vuln/cmd/govulncheck@v1.1.3

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=go.mod,target=./go.mod \
    --mount=type=bind,source=go.sum,target=./go.sum \
    go mod download

FROM golang:${GO_VERSION} AS builder

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=go.mod,target=./go.mod \
    --mount=type=bind,source=go.sum,target=./go.sum \
    go mod download

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=.,target=. \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /bin/pr2otel .


FROM gcr.io/distroless/static:nonroot

COPY --from=builder --chown=nonroot:nonroot /bin/pr2otel /bin/pr2otel

USER nonroot

ENTRYPOINT ["/bin/pr2otel"]