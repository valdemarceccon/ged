FROM golang:1.13-alpine as BUILD

WORKDIR /server

COPY . /server

RUN mkdir bin

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -ldflags="-w -s" -v -o /server/bin/main cmd/gedserver/main.go

FROM ubuntu:18.04

WORKDIR /app


COPY --from=BUILD /server/bin/main ./

ENTRYPOINT ["./main"]
