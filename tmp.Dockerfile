FROM alpine:3.14

RUN set -x \
    && apk add --update --no-cache \
       ca-certificates \
    && rm -rf /var/cache/apk/*
COPY gitshockd /usr/local/bin/

EXPOSE 8545 9632 1478
ENTRYPOINT ["gitshockd"]