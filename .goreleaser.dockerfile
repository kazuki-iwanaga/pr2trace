FROM gcr.io/distroless/static:nonroot

COPY --chown=nonroot:nonroot pr2otel /bin/pr2otel

USER nonroot

ENTRYPOINT ["/bin/pr2otel"]