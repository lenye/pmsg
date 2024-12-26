FROM gcr.io/distroless/static-debian12:latest-amd64
COPY pmsg /
ENTRYPOINT ["/pmsg"]