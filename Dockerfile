FROM gcr.io/distroless/static-debian11
COPY pmsg /
ENTRYPOINT ["/pmsg"]