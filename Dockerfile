FROM gcr.io/distroless/static-debian12
COPY pmsg /
ENTRYPOINT ["/pmsg"]