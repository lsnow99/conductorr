###
# Development Dockerfile for Conductorr
# Author Logan Snow
###

FROM golang:1.25

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD air -c .air.toml
