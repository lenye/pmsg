FROM scratch
COPY script/ca-certificates.crt /etc/ssl/certs/
COPY pmsg /
VOLUME ["/tmp"]
ENTRYPOINT ["/pmsg"]