# Build stage container
FROM golang:1.17.6-buster as builder

WORKDIR /app

COPY main.go go.mod go.sum ./
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

RUN go build -o kv-store-grpc

# Application container
FROM debian:buster-slim

ARG COMMAND
ENV COMMAND "$COMMAND"

COPY --from=builder /app/kv-store-grpc /app/kv-store-grpc

CMD /app/kv-store-grpc $COMMAND
