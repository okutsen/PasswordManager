FROM golang:1.18-alpine3.15 AS build

WORKDIR /github.com/okutsen/PasswordManager

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/main.go ./cmd/
COPY api/*.go ./api/
COPY pkg/*.go ./pkg/
COPY internal/*.go ./internal/

RUN cd ./cmd/ && go build -o ../build/out

FROM alpine:3.15 AS release

EXPOSE 10000

COPY --from=build /github.com/okutsen/PasswordManager/build/out /out

ENTRYPOINT [ "/out" ]

# Pass secrets to Docker
# golangci-lint