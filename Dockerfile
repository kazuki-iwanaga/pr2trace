ARG GO_VERSION=1.23

FROM golang:${GO_VERSION} AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/pr2otel .


FROM gcr.io/distroless/static:nonroot

COPY --from=builder --chown=nonroot:nonroot /bin/pr2otel /bin/pr2otel

USER nonroot

ENTRYPOINT ["/bin/pr2otel"]