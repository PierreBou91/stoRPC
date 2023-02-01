FROM golang:1.19 AS builder

COPY . /src

WORKDIR /src/storpc_server

RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM scratch

COPY --from=builder /src/storpc_server/server .

ENTRYPOINT ["./server"]