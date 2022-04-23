FROM alpine:latest

RUN apk update && apk upgrade && apk add --no-cache ca-certificates

COPY gomo /usr/bin/gomo

ENTRYPOINT ["/usr/bin/gomo"]
