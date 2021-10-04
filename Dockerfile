###
# Dockerfile for Conductorr
# Author Logan Snow
###

# Build the CSL web assembly module
# This stage of the build also copies the wasm_exec.js script to the web ui source
FROM golang:1.17 AS csl-build-env
COPY . /build
RUN apt-get update && apt-get install -y brotli
RUN cd /build && GOOS=js GOARCH=wasm go build -o dist/csl.wasm ./cmd/csl
RUN cd /build && brotli -f dist/csl.wasm
RUN cp $(go env GOROOT)/misc/wasm/wasm_exec.js /build/web/src/util

# Build the frontend web ui
FROM node:lts-alpine AS vue-build-env
COPY --from=csl-build-env /build/web .
RUN yarn install
RUN yarn build --emptyOutDir

# Build the full binary
FROM golang:1.17-alpine AS svr-build-env
RUN apk add build-base
RUN apk add git
RUN mkdir /src
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY --from=csl-build-env /build .
COPY --from=vue-build-env build/ web/build
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
    go build -o /build/conductorr \
    -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=$(git describe --tags)' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" \
    ./cmd/conductorr

# Copy the full binary into a minimal alpine image
FROM alpine
WORKDIR /app
COPY --from=svr-build-env /build/conductorr /app/conductorr
CMD ["/app/conductorr"]
