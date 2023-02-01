FROM golang:1.19 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
COPY storpc_server/server.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM scratch

COPY --from=builder /usr/src/app/server .

ENTRYPOINT server