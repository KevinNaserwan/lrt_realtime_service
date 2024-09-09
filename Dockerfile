FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./build/main ./cmd/main.go

# Path: config/Dockerfile.backend

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/build/main /app/main

COPY  --from=builder /app/log /app/log

ENTRYPOINT [ "/app/main" ]
