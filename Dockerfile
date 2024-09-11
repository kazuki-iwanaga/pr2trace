ARG GO_VERSION=1.22

FROM golang:${GO_VERSION} AS dev

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod \
    go install github.com/spf13/cobra-cli@1.3.0

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=go.mod,target=./go.mod \
    --mount=type=bind,source=go.sum,target=./go.sum \
    go mod download