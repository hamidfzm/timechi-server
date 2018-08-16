FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir /data
ADD data/config.yml.sample /data/config.yml
ADD timechi /
VOLUME /data

EXPOSE 8080

ENTRYPOINT ["/timechi"]
