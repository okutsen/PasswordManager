FROM golang:1.18-alpine3.15 AS build

WORKDIR /passwordmanager

COPY go.mod ./
COPY go.sum ./


COPY cmd/ ./cmd
COPY config/ ./config
COPY internal/ ./internal
COPY pkg/ ./pkg
COPY schema/ ./schema

RUN go build -o ./bin/pm ./cmd/


FROM alpine:3.15 AS release

COPY --from=build /passwordmanager/bin/pm /pm

ENV PM_PORT=10000
ENV PM_DB_HOST=postgres
ENV PM_DB_PORT=5432
ENV PM_DB_NAME=password_manager
ENV PM_DB_USERNAME=admin
ENV PM_DB_PASSWORD=12345
ENV PM_DB_SSL_MODE=disable

ENTRYPOINT ["/pm"]
