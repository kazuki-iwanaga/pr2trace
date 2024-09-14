ARG GO_VERSION=1.23

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