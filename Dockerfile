FROM golang:1.18-alpine3.15 AS build

WORKDIR /PasswordManager

COPY go.mod ./
COPY go.sum ./


COPY cmd/ ./cmd
COPY internal ./internal
COPY config ./config
COPY pkg/ ./pkg


RUN go build -o ./bin/pm ./cmd/


FROM alpine:3.15 AS release

COPY --from=build /PasswordManager/bin/pm /pm

ENV PM_PORT=10000

ENTRYPOINT ["/pm"]