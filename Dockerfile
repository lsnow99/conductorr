FROM golang:alpine AS go-build-env
ADD . /src
RUN cd /src && go build ./cmd/conductorr
RUN cd /src && go build ./cmd/migrations

FROM node:lts-alpine AS vue-build-env
COPY conductorr-web/ .
RUN npm install
RUN npm run build

FROM alpine
WORKDIR /app
COPY --from=go-build-env /src/conductorr /app/
COPY --from=go-build-env /src/migrations /app/
COPY --from=vue-build-env dist/ static/
ENTRYPOINT ./conductorr