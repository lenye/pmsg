FROM gcr.io/distroless/static-debian13:latest-amd64
COPY pmsg /
ENTRYPOINT ["/pmsg"]