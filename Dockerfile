###
# Dockerfile for Conductorr
# Author Logan Snow
###

ARG VERSION

# Build the CSL web assembly module
# This stage of the build also copies the wasm_exec.js script to the web ui source
FROM golang:1.17 AS csl-build-env
WORKDIR /build
RUN apt-get update && apt-get install -y brotli
COPY ./pkg ./pkg
COPY ./cmd/cslwasm ./cmd/cslwasm
COPY ./internal/csl ./internal/csl
COPY [ "go.mod", "go.sum", "./" ]
RUN GOOS=js GOARCH=wasm go build -o dist/csl.wasm ./cmd/cslwasm
RUN brotli -f dist/csl.wasm
RUN cp $(go env GOROOT)/misc/wasm/wasm_exec.js .

# Build the frontend web UI
FROM node:lts-alpine AS vue-build-env
WORKDIR /build
RUN npm install -g pnpm
COPY ./web ./web
COPY pnpm-lock.yaml .
COPY pnpm-workspace.yaml .
# Copy the exact wasm_exec.js file from the installation of Go that built the wasm module
COPY --from=csl-build-env /build/wasm_exec.js ./web/src/util
RUN cd web && pnpm install
RUN cd web && pnpm build

# Build the full binary
FROM golang:1.17-alpine AS svr-build-env
RUN apk add build-base
RUN apk add git
RUN mkdir /src
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./cmd/conductorr ./cmd/conductorr
COPY --from=csl-build-env /build .
COPY --from=vue-build-env /build/web/build web/build
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
    go build -o /build/conductorr \
    -ldflags="-s -w -X 'github.com/lsnow99/conductorr/settings.Version=${VERSION}' -X 'github.com/lsnow99/conductorr/settings.BuildMode=binary'" \
    ./cmd/conductorr

# Copy the full binary into a minimal alpine image
FROM alpine
WORKDIR /app
COPY --from=svr-build-env /build/conductorr /app/conductorr
CMD ["/app/conductorr"]
