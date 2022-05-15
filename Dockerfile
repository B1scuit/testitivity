FROM golang:1.17-alpine AS build

ADD ./ /go/app

WORKDIR /go/app
RUN apk update && apk upgrade && \
    apk add --no-cache bash git curl make gcc openssh

RUN make

#
# Get root CA's for google API
#
FROM alpine:latest as certHandler

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
#
# Create final image
#
FROM scratch

COPY --from=build /go/app/commander /opt/commander
COPY --from=certHandler /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /opt

ENTRYPOINT ["./commander"]