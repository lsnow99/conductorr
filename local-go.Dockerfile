###
# Development Dockerfile for Conductorr
# Author Logan Snow
###

FROM golang:1.23

WORKDIR /app

CMD go run ./cmd/conductorr
