FROM gcr.io/distroless/static:nonroot

COPY --chown=nonroot:nonroot pr2trace /bin/pr2trace

USER nonroot

ENTRYPOINT ["/bin/pr2trace"]