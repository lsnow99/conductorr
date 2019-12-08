FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build ./cmd/conductorr
RUN cd /src && go build ./cmd/migrations

FROM alpine
WORKDIR /app
COPY --from=build-env /src/conductorr /app/
COPY --from=build-env /src/migrations /app/
COPY conductorr-web/dist/ static/
ENTRYPOINT ./conductorr