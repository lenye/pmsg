FROM scratch
COPY /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY pmsg /
VOLUME ["/tmp"]
ENTRYPOINT ["/pmsg"]