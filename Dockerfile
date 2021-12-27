###
# Dockerfile for Conductorr
# Author Logan Snow
# Arguments:
# - VERSION can be set using $(git describe --tags)
###

# Build the CSL WebAssembly module
# This stage of the build also copies the wasm_exec.js script to the frontend source folder
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
# Build the shared library
COPY ./shared ./shared
RUN cd shared && pnpm install
RUN cd shared && pnpm build
# Build the frontend
COPY ./frontend ./frontend
COPY pnpm-lock.yaml .
COPY pnpm-workspace.yaml .
# Copy the exact wasm_exec.js file from the installation of Go that built the wasm module
COPY --from=csl-build-env /build/wasm_exec.js ./frontend/src/util
RUN cd frontend && pnpm install
RUN cd frontend && pnpm build

# Build the full binary
FROM golang:1.17-alpine AS svr-build-env
RUN apk add build-base
RUN apk add git
RUN mkdir /src
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY --from=csl-build-env /build/dist ./dist
COPY --from=vue-build-env /build/frontend/build frontend/build
COPY ./embed.go ./embed.go
COPY ./migrations ./migrations
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./cmd/conductorr ./cmd/conductorr
ARG VERSION=NULL
RUN CGO_ENABLED=0 GOOS=linux \
    go build -o /build/conductorr \
    -ldflags="-s -w -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.Version=${VERSION}' -X 'github.com/lsnow99/conductorr/internal/conductorr/settings.BuildMode=binary'" \
    ./cmd/conductorr

# Copy the full binary into a minimal alpine image
FROM alpine
WORKDIR /app
COPY --from=svr-build-env /build/conductorr /app/conductorr
CMD ["/app/conductorr"]
