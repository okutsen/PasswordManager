FROM golang:1.18-alpine3.15 AS build

WORKDIR /passwordmanager

COPY go.mod ./
COPY go.sum ./


COPY cmd/ ./cmd
COPY internal/ ./internal
COPY config/ ./config
COPY pkg/ ./pkg
COPY schema/ ./schema


RUN go build -o ./bin/pm ./cmd/


FROM alpine:3.15 AS release

COPY --from=build /passwordmanager/bin/pm /pm

ENV PM_PORT=10000

ENTRYPOINT ["/pm"]
