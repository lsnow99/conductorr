###
# Development Dockerfile for Conductorr
# Author Logan Snow
###

FROM golang:1.23

WORKDIR /app

RUN apt update && \
    apt install -y \
      entr

CMD find . -type f -regex ".*\.go" | entr -nr go run ./cmd/conductorr
